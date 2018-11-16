package solver

import (
	"../models"
)

// RemoveEmptyProductions should remove every expression that generates empty word
func RemoveEmptyProductions(automato models.Automato) models.Automato {
	nonEmptyProductions := []models.Expression{}
	emptyProductions := []models.Expression{}

	for _, expression := range automato.Expressions {
		for _, word := range expression.Words {
			if word.IsEmpty() {
				emptyProductions = append(emptyProductions, expression)
			}
		}
	}

	for _, expression := range emptyProductions {
		variable := expression.Variable
	}

	return models.Automato{
		Expressions: nonEmptyProductions,
	}
}
