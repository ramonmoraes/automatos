package main

import (
	"fmt"

	"./models"
	"./solver"
)

func main() {
	at := models.NewAutomato([]string{
		"S -> ASA | aB ",
		"A -> B | S",
		"B -> b | 0",
	})
	// at := models.NewAutomato([]string{"A -> aB | a", "B -> C | 0", "C -> D", "D -> ff"})
	fmt.Println("[Initial Automato]")
	at.Explain()

	chomAt := solver.Solve(at)
	// fmt.Println("[End Automato]")
	chomAt.Explain()
}
