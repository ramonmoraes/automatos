package solver

import (
	"../models"
)

// RemoveEmptyProductions should remove every expression that generates empty word
func RemoveEmptyProductions(automato models.Automato) models.Automato {
	nonEmptyProductions := make(map[models.Simbol][]models.Word)
	emptyProductions := make(map[models.Simbol][]models.Word)

	for variable, words := range automato.Expressions {
		isEmpty := false

		for _, word := range words {
			if word.IsEmpty() {
				isEmpty = true
			}
		}

		if isEmpty {
			emptyProductions[variable] = words
		} else {
			nonEmptyProductions[variable] = words
		}
	}

	atWithoutEmptyProductions := models.Automato{
		Expressions: nonEmptyProductions,
	}

	for emptyProductionVar, contentFromEmptyVar := range emptyProductions {
		for _, wordsToBeAppendedWithEmptyVarContent := range atWithoutEmptyProductions.Expressions {
			for _, word := range wordsToBeAppendedWithEmptyVarContent {
				if word.Contains(emptyProductionVar) {
					nw := word.ReplaceWords(emptyProductionVar, contentFromEmptyVar)
					wordsToBeAppendedWithEmptyVarContent = append(wordsToBeAppendedWithEmptyVarContent, nw...)
				}
			}
		}
	}

	return atWithoutEmptyProductions
}
