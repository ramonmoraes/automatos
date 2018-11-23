package solver

import (
	"testing"

	"../models"
)

func TestRealocateTerminal(t *testing.T) {
	at := models.NewAutomato([]string{"A -> Ff"})
	at = RealocateTerminals(at)
	if len(at.Expressions) != 2 {
		t.Error("Should've generated two expressions")
	}

	at = models.NewAutomato([]string{
		"A -> Cf",
		"C -> hh",
	})
	at = RealocateTerminals(at)
	if len(at.Expressions) != 4 {
		t.Error("Should've generated four expressions")
	}
}

func TestGenerateNewSimbol(t *testing.T) {
	at := models.NewAutomato([]string{"A -> Ff"})

	s, _ := models.NewSimbol("f")
	generatedSimbol := GenerateNewSimbol(at.GetVariableList(), s)
	if generatedSimbol.Value == "F" {
		t.Error("Should not generate already existing simbol")
	}

	at = models.NewAutomato([]string{"A -> Fc"})
	s, _ = models.NewSimbol("c")
	generatedSimbol = GenerateNewSimbol(at.GetVariableList(), s)
	if generatedSimbol.Value != "C" {
		t.Error("Should generate C as simbol")
	}
}
