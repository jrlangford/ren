package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

var debug = flag.Bool("debug", false, "Run ren in debug mode")
var inputData = flag.String("c", "", "Set the csv input string")
var templateFile = flag.String("t", "", "Set the template input file")

func printErr(e error, s string) {
	fmt.Fprintf(os.Stderr, "Error: %s -> %s\n", s, e)
}

func csvKeyValuesToMap(s string, dataMap map[string]string) (err error) {
	csvParsed := csv.NewReader(strings.NewReader(s))
	for {
		record, err := csvParsed.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			printErr(err, "CSV Parsing")
			return err
		}

		for _, entry := range record {
			substrings := strings.Split(entry, ":")
			if len(substrings) != 2 {
				err = errors.New(fmt.Sprintf("Unexpected record format: %s, Data: %v", record, substrings))
				printErr(err, "CSV Format")
				return err
			}

			key := strings.Trim(substrings[0], " ")
			value := strings.Trim(substrings[1], " ")

			dataMap[key] = value
		}

		if *debug {
			for key, value := range dataMap {
				fmt.Fprintf(os.Stderr, "[%s]:%s\n", key, value)
			}
		}
	}
	return err
}

func renderTemplate(tmplFilename, inputTuples string, wr io.Writer) (err error) {
	tmpl, err := template.ParseFiles(tmplFilename)
	if err != nil {
		printErr(err, "Could not parse template")
		return err
	}

	var dataMap map[string]string
	dataMap = make(map[string]string)

	err = csvKeyValuesToMap(inputTuples, dataMap)
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(wr, path.Base(tmplFilename), dataMap)
	if err != nil {
		printErr(err, "Could not execute template")
		return err
	}
	return err
}

func main() {
	flag.Parse()

	if *inputData == "" {
		fmt.Fprintf(os.Stderr, "CSV data parameter missing, aborting\n")
		os.Exit(1)
	}

	if *templateFile == "" {
		fmt.Fprintf(os.Stderr, "Template file parameter missing, aborting\n")
		os.Exit(1)
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "Input string: \n%s\n", *inputData)
	}

	err := renderTemplate(*templateFile, *inputData, os.Stdout)
	if err != nil {
		os.Exit(1)
	}
}
