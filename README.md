#REN

A file renderer written in go based on the ["text/template library"](http://golang.org/pkg/text/template/)

##Build
```
$ go build
```

##Install
```
$ go install
```

##Usage

REN takes a json string and a template file as inputs and outputs the processed string to stdout.

```
$ ren -c <input_csv_string> -t <input_template_file>

```
###Example

input_template (test_data/config.in):
```
App config file
app_host: {{ .host }}
app_port: {{ .port }}
```

Rendering:
```
$ ren -c "host:localhost,port:8080" -t test_data/config_file
```

Output:
```
App config file
app_host: localhost
app_port: 8080
```
