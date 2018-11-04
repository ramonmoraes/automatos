package solver

import (
	"../models"
)

func UselessSimbol(automato models.Automato) models.Automato {
	validExpressionList := []models.Expression{}
	for _, expr := range automato.Expressions {
		for _, word := range expr.Generated {
			if word.ContainTerminal() {
				validExpressionList = append(validExpressionList, expr)
			}
		}
	}

	return models.Automato{
		Expressions: validExpressionList,
	}
}
