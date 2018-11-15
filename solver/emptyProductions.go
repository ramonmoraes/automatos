package solver

import (
	"../models"
)

// EmptyProductions should remove every expression that generates empty word
func EmptyProductions(automato models.Automato) models.Automato {
	nonEmptyProductions := []models.Expression{}

	return models.Automato{
		Expressions: nonEmptyProductions,
	}
}
