package models

import (
	"fmt"
	"log"
)

// Automato should be a list of expressions
// It can be created with NewAutomato method or normally
type Automato struct {
	Expressions []Expression
}

// NewAutomato should create an Automato parsing the string equations given
func NewAutomato(expressionStringList []string) Automato {
	expressionList := []Expression{}
	for _, expressionString := range expressionStringList {
		expression, err := NewExpression(expressionString)
		if err != nil {
			log.Fatal(err)
		}

		expressionList = append(expressionList, expression)
	}
	return Automato{
		Expressions: expressionList,
	}
}

// Explain should call explain for every expression in automato
func (a *Automato) Explain() {
	fmt.Println("Explaining Automato:")
	for _, expression := range a.Expressions {
		expression.Explain()
	}
}
