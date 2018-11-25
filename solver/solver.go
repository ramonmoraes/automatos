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
	at := RemoveEmptyProductions(a) // null productions
	// fmt.Println("[After Removing Empty Productions]")
	// at.Explain()

	at = UselessSimbol(at) // unit productions
	// fmt.Println("[After removing useless simbol]")
	// at.Explain()

	// at = RealocateTerminals(at)
	// fmt.Println("[After Realocating Terminals]")
	// at.Explain()

	// at = RealocateVariables(at)

	return models.Fix(at)
}
