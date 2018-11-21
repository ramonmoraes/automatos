package solver

import (
	"fmt"

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
		for expressionSimbol, wordsToBeAppendedWithEmptyVarContent := range atWithoutEmptyProductions.Expressions {
			for i, word := range wordsToBeAppendedWithEmptyVarContent {
				if word.Contains(emptyProductionVar) {
					newWords := word.ReplaceWords(emptyProductionVar, contentFromEmptyVar)
					nonEmptyWords := []models.Word{}
					for _, w := range newWords {
						if !w.IsEmpty() {
							nonEmptyWords = append(nonEmptyWords, w)
						}
					}

					newWordsToBeAppended := []models.Word{}
					newWordsToBeAppended = append(newWordsToBeAppended, wordsToBeAppendedWithEmptyVarContent[:i]...)
					newWordsToBeAppended = append(newWordsToBeAppended, wordsToBeAppendedWithEmptyVarContent[i+1:]...)
					newWordsToBeAppended = append(newWordsToBeAppended, nonEmptyWords...)
					atWithoutEmptyProductions.Expressions[expressionSimbol] = newWordsToBeAppended
					i = i + len(nonEmptyWords) - 1
				}
			}
		}
	}

	return atWithoutEmptyProductions
}

func explain(words []models.Word) {
	// fmt.Println("Explaining list of words")
	for _, w := range words {
		fmt.Println(w.ToString())
	}
}
