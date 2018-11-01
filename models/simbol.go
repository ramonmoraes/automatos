package models

import (
	"errors"
	"fmt"
	"strings"
)

// Simbol should be the base struct of the program
// but it should always be initiated with the NewSimbol func
type Simbol struct {
	Value      string
	IsVariable bool
}

// NewSimbol should return a new instance of simbol
func NewSimbol(value string) (simbol, error) {
	if len(value) != 1 {
		errorMessage := fmt.Sprintf("A simbol '%s' can not have a length different from 1", value)
		return simbol{}, errors.New(errorMessage)
	}

	return simbol{
		Value:      value,
		IsVariable: value == strings.ToUpper(value),
	}, nil
}
