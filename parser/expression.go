package parser

import (
	"fmt"
	"strings"

	"../models"
)

// Expression should be our base struct for using in expressions
// But it should be created with NewExpression func
type Expression struct {
	Creator   models.Simbol
	Generated []models.Word
}

// NewExpression should receive a string and return a Expression expression
// of which simbol generates which
func NewExpression(expression string) (Expression, error) {
	splited := strings.Split(expression, "->")
	creator, err := models.NewSimbol(strings.TrimSpace(splited[0]))

	if err != nil {
		return Expression{}, err
	}

	generatedList := strings.Split(strings.TrimSpace(splited[1]), "|")
	wordList := []models.Word{}

	for _, word := range generatedList {
		wordList = append(wordList, models.NewWord(word))
	}

	return Expression{
		Creator:   creator,
		Generated: wordList,
	}, nil
}

// Explain should only print a information of what the expression resulted
func (e *Expression) Explain() {
	words := []string{}
	for _, word := range e.Generated {
		words = append(words, word.ToString())
	}
	fmt.Println("Explaining:")
	fmt.Printf(
		"The variable %s, may generate the word(s) '%s'\n",
		e.Creator.Value, strings.Join(words, " or "),
	)
}
