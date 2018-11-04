package main

import (
	"fmt"

	"./models"
	"./solver"
)

func main() {
	at := models.NewAutomato([]string{"A -> aB | a", "B -> E", "E -> c | 0"})
	sim := at.GetPossibleCreatedsSimbols()
	fmt.Println(sim)

	anotherAt := solver.UselessSimbol(at)
	fmt.Println(anotherAt)
	// valList := []string{}
	// for _, simbol := range sim {
	// 	valList = append(valList, simbol.Value)
	// }
	// fmt.Println(valList)
}
