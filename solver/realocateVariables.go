package solver

import (
	"strings"

	"../models"
)

// RealocateVariables should return an automato without words with 3 or more variables
func RealocateVariables(a models.Automato) models.Automato {
	for _, words := range a.Expressions {
		for _, word := range words {
			if len(word.GetVariables()) > 2 {
				return splitWords(a, word)
			}
		}
	}
	return a
}

func splitWords(a models.Automato, word models.Word) models.Automato {
	var wordPair []string
	var aux string

	for i, char := range strings.Split(word.ToString(), "") {
		aux = aux + char
		if i%2 == 0 {
			wordPair = append(wordPair, aux)
			aux = ""
		}
	}

	var replacedAt models.Automato

	for _, s := range wordPair {
		optimalSimbol, _ := models.NewSimbol(strings.Split(s, "")[0])
		newSimbol := GenerateNewSimbol(a.GetVariableList(), optimalSimbol)
		replacedAt = replaceStringsBigger3(a, s, newSimbol.Value)
		newWord := models.NewWord(s)
		replacedAt.Expressions[newSimbol] = []models.Word{newWord}
	}

	return RealocateVariables(replacedAt)
}

func replaceStringsBigger3(a models.Automato, toBeReplacedString string, newString string) models.Automato {
	replacedAt := models.Automato{
		Expressions: make(map[models.Simbol][]models.Word),
	}

	for simbol, words := range a.Expressions {
		for _, word := range words {
			wordString := word.ToString()
			if len(word.Simbols) > 2 {
				wordString = strings.Replace(wordString, toBeReplacedString, newString, -1)
			}
			newWord := models.NewWord(wordString)
			_, ok := replacedAt.Expressions[simbol]
			if ok {
				replacedAt.Expressions[simbol] = append(replacedAt.Expressions[simbol], newWord)
			} else {
				replacedAt.Expressions[simbol] = []models.Word{newWord}
			}
		}
	}

	return replacedAt
}
