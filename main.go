package main

import (
	"log"

	"./parser"
)

func main() {
	ex, err := parser.NewExpression("A -> aB | a")
	if err != nil {
		log.Fatal(err)
	}
	ex.Explain()
}
