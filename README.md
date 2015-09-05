#REN

A file renderer written in go based on the following libraries:
* ["text/template"](http://golang.org/pkg/text/template/)
* ["jeffail/gabs"](https://github.com/jeffail/gabs)

##Build
```
$ go get
$ go build
```

##Install
```
$ go install
```

##Usage

REN takes a json file and a template as inputs and outputs the processed string to stdout.

```
$ ren -j <input_json_file> -t <input_template_file>

```
###Example

input_json_file (example/data.json):
```
{
	"host": "localhost",
	"port": 8080
}
```

input_template_file (example/config.in):
```
App config file
app_host: {{ .host }}
app_port: {{ .port }}
```

Rendering:
```
$ ren -j example/data.json -t example/config.in
```

Output:
```
App config file
app_host: localhost
app_port: 8080
```
