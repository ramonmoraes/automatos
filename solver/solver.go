package solver

import "../models"

// Solve should just return the chomsky form
func Solve(a models.Automato) models.Automato {
	return Chomsky(a)
}

// Chomsky should return a simplified automato
func Chomsky(a models.Automato) models.Automato {
	at := RemoveEmptyProductions(a)
	at = UselessSimbol(at)
	return at
}
