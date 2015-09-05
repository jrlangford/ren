package main

import (
	"flag"
	"fmt"
	"github.com/jeffail/gabs"
	"os"
	"path"
	"text/template"
)

var debug = flag.Bool("debug", false, "Run ren in debug mode")
var jsonData = flag.String("j", "", "Set the json input string")
var templateFile = flag.String("t", "", "Set the template input file")

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %s -> %s\n", s, e)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	if *jsonData == "" {
		fmt.Fprintf(os.Stderr, "JSon data parameter missing, aborting\n")
		os.Exit(1)
	}

	if *templateFile == "" {
		fmt.Fprintf(os.Stderr, "Template file parameter missing, aborting\n")
		os.Exit(1)
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "jsonData: \n%s\n", jsonData)
	}

	templates, err := template.ParseFiles(*templateFile)
	check(err, "Could not parse template")

	jsonParsed, err := gabs.ParseJSON([]byte(*jsonData))
	check(err, "Could not parse json")

	children, err := jsonParsed.ChildrenMap()
	check(err, "Could not read children")

	if *debug {
		for key, child := range children {
			fmt.Fprintf(os.Stderr, "key: %v, value: %v\n", key, child.Data())
		}
	}

	var typedMap map[string]string
	typedMap = make(map[string]string)
	for key, child := range children {
		typedMap[key] = fmt.Sprintf("%v", child.Data())
	}

	err = templates.ExecuteTemplate(os.Stdout, path.Base(*templateFile), typedMap)
	check(err, "Could not execute template")

}
