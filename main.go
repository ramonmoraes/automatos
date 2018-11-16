package main

import (
	"./models"
)

func main() {
	// at := models.NewAutomato([]string{"A -> aB | a", "B -> E", "E -> c | 0"})
	at := models.NewAutomato([]string{"A -> aB | a", "B -> 0"})
	// sim := at.GetPossibleCreatedsSimbols()
	// fmt.Println(sim)
	at.Explain()
	// anotherAt := solver.UselessSimbol(at)
	// fmt.Println(anotherAt)
	// valList := []string{}
	// for _, simbol := range sim {
	// 	valList = append(valList, simbol.Value)
	// }
	// fmt.Println(valList)
}
