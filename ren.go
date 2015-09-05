package main

import (
	"flag"
	"fmt"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"os"
	"text/template"
	"path"
)

var debug = flag.Bool("debug", false, "Run ren in debug mode")
var dataFile = flag.String("j", "", "Set the json input file")
var templateFile = flag.String("t", "", "Set the template input file")

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %s -> %s\n", s, e)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	if *dataFile == "" {
		fmt.Fprintf(os.Stderr, "Data file parameter missing, aborting\n")
		os.Exit(1)
	}

	if *templateFile == "" {
		fmt.Fprintf(os.Stderr, "Template file parameter missing, aborting\n")
		os.Exit(1)
	}

	jsonString, err := ioutil.ReadFile(*dataFile)
	check(err, "Could not read file")

	if *debug {
		fmt.Fprintf(os.Stderr, "jsonString: \n%s\n", jsonString)
	}

	templates, err := template.ParseFiles(*templateFile)
	check(err, "Could not parse template")

	jsonParsed, err := gabs.ParseJSON(jsonString)
	check(err, "Could not parse json")

	children, err := jsonParsed.ChildrenMap()
	check(err, "Could not read children")

	if *debug {
		for key, child := range children {
			fmt.Fprintf(os.Stderr, "key: %v, value: %v\n", key, child.Data())
		}
	}

	err = templates.ExecuteTemplate(os.Stdout, path.Base(*templateFile), children)
	check(err, "Could not execute template")

}
