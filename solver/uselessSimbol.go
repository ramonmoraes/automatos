package solver

import (
	"../models"
)

// UselessSimbol or unitProduction
// UselessSimbol should remove every useless simbol
func UselessSimbol(automato models.Automato) models.Automato {
	newAutomato := automato
	for originalSimbol, words := range removeSingleRecursive(automato).Expressions {
		var wordList []models.Word
		for _, word := range words {
			if len(word.Simbols) == 1 && word.Simbols[0].IsVariable {
				nw := word.ReplaceWords(
					word.Simbols[0],
					automato.Expressions[word.Simbols[0]],
				)
				wordList = append(wordList, nw...)
			} else {
				wordList = append(wordList, word)
			}
		}
		newAutomato.Expressions[originalSimbol] = wordList
	}
	return newAutomato
}

func removeSingleRecursive(automato models.Automato) models.Automato {
	at := automato
	for simbol, words := range at.Expressions {
		var nw []models.Word
		for _, word := range words {
			if len(word.Simbols) != 1 || word.Simbols[0] != simbol {
				nw = append(nw, word)
			}
		}
		at.Expressions[simbol] = nw
	}
	return at
}
