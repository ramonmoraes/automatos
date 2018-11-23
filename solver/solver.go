package solver

import (
	"../models"
)

// Solve should just return the chomsky form
func Solve(a models.Automato) models.Automato {
	return Chomsky(a)
}

// Chomsky should return a simplified automato
func Chomsky(a models.Automato) models.Automato {
	at := RemoveEmptyProductions(a)
	// fmt.Println("[After Removing Empty Productions]")
	// at.Explain()

	at = UselessSimbol(at)
	// fmt.Println("[After removing useless simbol]")
	// at.Explain()

	at = RealocateTerminals(at) // broken <--
	// fmt.Println("[After Realocating Terminals]")

	// at.Explain()

	return at
}
