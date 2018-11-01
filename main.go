package main

import (
	"fmt"
	"log"

	"./parser"
)

func main() {
	ex, err := parser.NewTrasformation("A -> ab")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ex)
}
