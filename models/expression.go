package models

import (
	"fmt"
	"strings"
)

// Expression should be our base struct for using in expressions
// But it should be created with NewExpression func
type Expression struct {
	Variable Simbol
	Words    []Word
}

// NewExpression should receive a string and return a Expression expression
// of which simbol generates which
func NewExpression(expression string) (Expression, error) {
	splited := strings.Split(expression, "->")
	creator, err := NewSimbol(strings.TrimSpace(splited[0]))

	if err != nil {
		return Expression{}, err
	}

	generatedList := strings.Split(strings.TrimSpace(splited[1]), "|")
	wordList := []Word{}

	for _, word := range generatedList {
		wordList = append(wordList, NewWord(word))
	}

	return Expression{
		Variable: creator,
		Words:    wordList,
	}, nil
}

// Explain should only print a information of what the expression resulted
func (e *Expression) Explain() {
	words := []string{}
	for _, word := range e.Words {
		words = append(words, word.ToString())
	}
	fmt.Printf(
		"The variable %s, may generate the word(s) '%s'\n",
		e.Variable.Value, strings.Join(words, " or "),
	)
}
