package parser

import (
	"strings"

	"../models"
)

// Transformation should be our base struct for using in expressions
// But it should be created with NewTransformation func
type Transformation struct {
	Creator   models.Simbol
	Generated []models.Simbol
}

// NewTrasformation should receive a string and return a Transformation expression
// of which simbol generates which
func NewTrasformation(expression string) (Transformation, error) {
	splited := strings.Split(expression, "->")
	creator, err := models.NewSimbol(strings.TrimSpace(splited[0]))

	if err != nil {
		return Transformation{}, err
	}

	generatedList := strings.Split(strings.TrimSpace(splited[1]), "")
	simbolList := []models.Simbol{}

	for _, char := range generatedList {
		simbol, err := models.NewSimbol(char)

		if err != nil {
			return Transformation{}, err
		}

		simbolList = append(simbolList, simbol)
	}

	return Transformation{
		Creator:   creator,
		Generated: simbolList,
	}, nil
}
