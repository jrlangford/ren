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
$ ren -j <input_json> -t <input_template>

```
###Example
```
$ ren -j example/data.json -t example/config.in
```
