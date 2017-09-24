package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

type options struct {
	all  *bool
	path []string
}

func parseArguments() options {
	initialBool := true
	options := options{&initialBool, nil}
	options.all = flag.Bool("a", false, "do not ignore entries starting with .")
	flag.Parse()
	options.path = flag.Args()
	if len(options.path) == 0 {
		options.path = []string{"./"}
	}
	return options
}

func ls(options options, file string) int {
	files, err := ioutil.ReadDir(file)
	if err != nil {
		return -1
	}
	for _, f := range files {
		fileName := f.Name()
		if !(fileName[0] == '.' && !*options.all) {
			fmt.Println(fileName)
		}
	}
	return 0
}

func main() {
	options := parseArguments()
	for _, file := range options.path {
		ls(options, file)
	}
}
