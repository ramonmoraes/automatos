package main

import (
	"fmt"

	"./models"
	"./solver"
)

func main() {
	print("\033[H\033[2J") // clean console
	// at := models.NewAutomato([]string{
	// 	"S -> ASA | aB ",
	// 	"A -> B | S",
	// 	"B -> b | 0",
	// })

	at := models.NewAutomato([]string{
		"A -> xB | B",
		"B -> fC | 0",
		"C -> C | g | 0",
	})

	fmt.Println("[Initial Automato]")
	at.Explain()

	chomAt := solver.Solve(at, false)
	fmt.Println("[End Automato]")
	chomAt.Explain()
}
