package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func parseArguments() []string {
	flag.Parse()
	arguments := flag.Args()
	if len(arguments) == 0 {
		arguments = []string{"./"}
	}
	return arguments
}

func ls(file string) int {
	files, err := ioutil.ReadDir(file)
	if err != nil {
		return -1
	}
	for _, f := range files {
		fileName := f.Name()
		if fileName[0] != '.' {
			fmt.Println(fileName)
		}
	}
	return 0
}

func main() {
	arguments := parseArguments()
	for _, file := range arguments {
		ls(file)
	}
}
