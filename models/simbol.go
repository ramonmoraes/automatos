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
	IsEmpty    bool
}

// NewSimbol should return a new instance of simbol
func NewSimbol(value string) (Simbol, error) {
	if len(value) != 1 {
		errorMessage := fmt.Sprintf("A Simbol '%s' can not have a length different from 1", value)
		return Simbol{}, errors.New(errorMessage)
	}

	if value == "0" {
		return Simbol{
			Value:      value,
			IsVariable: false,
			IsEmpty:    true,
		}, nil
	}

	return Simbol{
		Value:      value,
		IsVariable: value == strings.ToUpper(value),
	}, nil
}
