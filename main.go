package main

import (
	"fmt"

	"./models"
	"./solver"
)

func main() {
	// at := models.NewAutomato([]string{
	// 	"S -> ASA | aB ",
	// 	"A -> B | S",
	// 	"B -> b | 0",
	// })

	at := models.NewAutomato([]string{
		"S -> aX | bY | b",
		"X -> x | 0 | d",
		"Y -> y | 0",
	})

	fmt.Println("[Initial Automato]")
	at.Explain()

	chomAt := solver.Solve(at, false)
	fmt.Println("[End Automato]")
	chomAt.Explain()
}
