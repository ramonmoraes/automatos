package solver

import (
	"fmt"

	"../models"
)

// Solve should just return the chomsky form
func Solve(a models.Automato, verbose bool) models.Automato {
	return Chomsky(a, verbose)
}

// Chomsky should return a simplified automato
func Chomsky(a models.Automato, verbose bool) models.Automato {
	at := RemoveEmptyProductions(removeSingleRecursive(a)) // null productions
	if verbose {
		fmt.Println("[After Removing Empty/Null Productions]")
		at.Explain()
	}

	at = UselessSimbol(at) // unit productions
	if verbose {
		fmt.Println("[After removing useless simbol (unit productions)]")
		at.Explain()
	}
	at = RealocateTerminals(at)

	if verbose {
		fmt.Println("[After Realocating Terminals]")
		at.Explain()
	}

	at = RealocateVariables(at)
	if verbose {
		fmt.Println("[After Realocating Variables]")
		at.Explain()
	}
	return at
}
