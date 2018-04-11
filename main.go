package main

import (
	"fmt"
	"github.com/andreluzz/cd-test/parser"
)

var version = "Development"

func main() {
	parsedString := parser.Translate("Testing continuous delivery features.")
	fmt.Println(parsedString)
	fmt.Println("Version: " + version)
}