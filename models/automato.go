package models

import (
	"fmt"
	"log"
	"strings"
)

// Automato should be a list of expressions
// It can be created with NewAutomato method or normally
type Automato struct {
	Expressions map[Simbol][]Word
}

// NewAutomato should create an Automato parsing the string equations given
func NewAutomato(expressionStringList []string) Automato {
	expressionMap := make(map[Simbol][]Word)
	for _, expressionString := range expressionStringList {
		simbol, word := newExpression(expressionString)
		expressionMap[simbol] = word
	}

	return Automato{
		Expressions: expressionMap,
	}
}

func newExpression(expression string) (Simbol, []Word) {
	splited := strings.Split(expression, "->")
	creator, err := NewSimbol(strings.TrimSpace(splited[0]))

	if err != nil {
		log.Fatal(err)
	}

	generatedList := strings.Split(strings.TrimSpace(splited[1]), "|")
	wordList := []Word{}

	for _, word := range generatedList {
		wordList = append(wordList, NewWord(word))
	}

	return creator, wordList
}

// Explain should describe each automato's expression
func (a *Automato) Explain() {
	fmt.Println("Explaining automato:")
	for simbol, words := range a.Expressions {
		explainedWords := []string{}
		for _, word := range words {
			explainedWords = append(explainedWords, word.ToString())
		}
		fmt.Printf("%s -> %s\n", simbol.Value, strings.Join(explainedWords, " | "))
	}
}
