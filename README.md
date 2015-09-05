#REN

A file renderer written in go based on the following libraries:
* ["text/template"](http://golang.org/pkg/text/template/)
* ["jeffail/gabs"](https://github.com/jeffail/gabs)

##Usage

REN takes a json file and a template as inputs and outputs the processed string to stdout.

```
$ ./ren -j data.json -t myTemplate.in
```
