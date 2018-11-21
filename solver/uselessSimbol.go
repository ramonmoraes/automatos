package solver

import (
	"../models"
)

// UselessSimbol should remove every useless simbol
func UselessSimbol(automato models.Automato) models.Automato {
	newAutomato := automato
	for _, words := range automato.Expressions {
		if len(words) == 1 {
			word := words[0]
			simbolToBeReplaced := word.Simbols[0]
			if len(word.Simbols) == 1 && simbolToBeReplaced.IsVariable {
				replacedAt := replaceSimbol(
					newAutomato,
					simbolToBeReplaced,
					automato.Expressions[simbolToBeReplaced],
				)
				return UselessSimbol(replacedAt)
			}
		}
	}
	return newAutomato
}

func replaceSimbol(automato models.Automato, initialSimbol models.Simbol, newWord []models.Word) models.Automato {
	newAt := automato
	for simbol, words := range automato.Expressions {
		for _, word := range words {
			if word.Contains(initialSimbol) {
				newWords := word.ReplaceWords(initialSimbol, newWord)
				newAt.Expressions[simbol] = newWords
			}
		}
	}
	delete(newAt.Expressions, initialSimbol)
	return newAt
}
