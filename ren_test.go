package main

import (
	"bytes"
	"fmt"
	"testing"
)

func failOnStringMismatch(expectedString, outputString string, t *testing.T) {
	if outputString != expectedString {
		fmt.Printf("--------------------\n")
		fmt.Printf("Expected String:\n")
		fmt.Printf("%s\n", expectedString)
		fmt.Printf("--------------------\n")
		fmt.Printf("Output String:\n")
		fmt.Printf("%s\n", outputString)
		t.Fail()
	}
}

func init() {
	fmt.Println("Starting tests.\nError messages are expected as part of some successful tests.")
}

func TestTupleIsMapped(t *testing.T) {
	s := "host:localhost,port:8080"

	var dataMap map[string]string
	dataMap = make(map[string]string)

	err := csvKeyValuesToMap(s, dataMap)
	if err != nil {
		t.Fatalf(err.Error())
	}

	failOnStringMismatch("localhost", dataMap["host"], t)
	failOnStringMismatch("8080", dataMap["port"], t)
}

func TestWhitespaceIsTrimmed(t *testing.T) {
	s := "     host:localhost,     port:      8080"

	var dataMap map[string]string
	dataMap = make(map[string]string)

	err := csvKeyValuesToMap(s, dataMap)
	if err != nil {
		t.Fatalf(err.Error())
	}

	failOnStringMismatch("localhost", dataMap["host"], t)
	failOnStringMismatch("8080", dataMap["port"], t)
}

func TestMultilineTuplesAreMapped(t *testing.T) {
	s := "host:localhost\nport:8080"

	var dataMap map[string]string
	dataMap = make(map[string]string)

	err := csvKeyValuesToMap(s, dataMap)
	if err != nil {
		t.Fatalf(err.Error())
	}

	failOnStringMismatch("localhost", dataMap["host"], t)
	failOnStringMismatch("8080", dataMap["port"], t)
}

func TestInvalidTupleFails(t *testing.T) {
	s := "host:localhost,port8080"

	var dataMap map[string]string
	dataMap = make(map[string]string)

	err := csvKeyValuesToMap(s, dataMap)

	expectedString := "Unexpected record format: [host:localhost port8080], Data: [port8080]"
	outputString := fmt.Sprintf("%s", err)

	failOnStringMismatch(expectedString, outputString, t)
}

func TestValidInputIsRendered(t *testing.T) {
	s := "host:localhost,port:8080"
	tmplFilename := "test_data/config_file"
	var outputBytes bytes.Buffer

	expectedString := `App config file
app_host: localhost
app_port: 8080
`

	err := renderTemplate(tmplFilename, s, &outputBytes)
	if err != nil {
		t.Fatalf(err.Error())
	}

	failOnStringMismatch(expectedString, outputBytes.String(), t)

}

func TestFailureOnMissingTemplate(t *testing.T) {
	s := "host:localhost,port:8080"
	tmplFilename := "test_data/non_existent_file"
	var outputBytes bytes.Buffer

	err := renderTemplate(tmplFilename, s, &outputBytes)

	expectedString := "open test_data/non_existent_file: no such file or directory"
	outputString := fmt.Sprintf("%s", err)

	failOnStringMismatch(expectedString, outputString, t)

}
