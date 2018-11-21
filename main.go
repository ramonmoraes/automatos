package main

import (
	"fmt"

	"./models"
	"./solver"
)

func main() {
	// at := models.NewAutomato([]string{"A -> aB | a", "B -> E", "E -> c | 0"})
	at := models.NewAutomato([]string{"A -> aB | a", "B -> C | 0", "C -> D", "D -> ff"})
	fmt.Println("[Start]")
	at.Explain()

	chomAt := solver.Solve(at)
	fmt.Println("[End]")
	chomAt.Explain()
}
