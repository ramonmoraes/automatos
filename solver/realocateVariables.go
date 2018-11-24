package solver

import (
	"fmt"
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
		replacedAt = replaceStrings(a, s, newSimbol.Value)
		newWord := models.NewWord(s)
		replacedAt.Expressions[newSimbol] = []models.Word{newWord}
		fmt.Printf("%s\n", s)
	}
	return RealocateVariables(replacedAt)
}
