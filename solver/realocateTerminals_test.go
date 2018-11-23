package solver

import (
	"testing"

	"../models"
)

func TestRealocateTerminal(t *testing.T) {
	at := models.NewAutomato([]string{"A -> Ff"})
	at = RealocateTerminals(at)
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
