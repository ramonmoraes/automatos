package main

import (
	"log"

	"./models"
)

func main() {
	ex, err := models.NewExpression("A -> aB | a")
	if err != nil {
		log.Fatal(err)
	}
	ex.Explain()
}
