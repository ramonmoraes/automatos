package simbol

import (
	"errors"
	"strings"
)

type simbol struct {
	Value      string
	IsVariable bool
}

// New should return a new instance of simbol
func New(value string) (simbol, error) {
	if len(value) != 1 {
		return simbol{}, errors.New("A simbol can not have a length different from 1")
	}

	return simbol{
		Value:      value,
		IsVariable: value == strings.ToUpper(value),
	}, nil
}
