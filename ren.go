package main

import (
	"encoding/csv"
	"flag"
	"io"
	"strings"
	"fmt"
	"os"
	"path"
	"text/template"
)

var debug = flag.Bool("debug", false, "Run ren in debug mode")
var inputData = flag.String("c", "", "Set the csv input string")
var templateFile = flag.String("t", "", "Set the template input file")

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %s -> %s\n", s, e)
		os.Exit(1)
	}
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

	tmpl, err := template.ParseFiles(*templateFile)
	check(err, "Could not parse template")

	csvParsed := csv.NewReader(strings.NewReader(*inputData))

	var dataMap map[string]string
	dataMap = make(map[string]string)

	for {
		record, err := csvParsed.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			check(err, "CSV Parsing")
		}
		if *debug {
			for index, entry := range record {
				fmt.Fprintf(os.Stderr, "%d:%s\n", index, entry)
			}
		}
		
		for _, entry := range record {
			substrings := strings.Split(entry, ":")
			if len(substrings) != 2 {
				fmt.Fprintf(os.Stderr, "Unexpected record format: %s, Data: %v, aborting\n", record, substrings)
				os.Exit(1)
			}
			
			key := strings.Trim(substrings[0], " ")
			value := strings.Trim(substrings[1], " ")
			
			dataMap[key] = value
		}
	}

	err = tmpl.ExecuteTemplate(os.Stdout, path.Base(*templateFile), dataMap)
	check(err, "Could not execute template")

}
