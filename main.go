package main

import (
	"fmt"
	"github.com/andreluzz/cd-test/parser"
)

func main() {
	parsedString := parser.Translate("Testing continuous delivery features.")
	fmt.Println(parsedString)
}