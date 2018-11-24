package solver

import (
	"fmt"
	"testing"

	"../models"
)

func TestRealocateTerminal(t *testing.T) {
	at := models.NewAutomato([]string{"A -> Ff"})
	at = RealocateTerminals(at)
	if len(at.Expressions) != 2 {
		t.Error("Should've generated two expressions")
		at.Explain()
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

func TestReplaceFirstTerminalForNewSimbol(t *testing.T) {
	at := models.NewAutomato([]string{"A -> Ff | c"})
	for _, words := range at.Expressions {
		at = ReplaceFirstTerminalForNewSimbol(at, words[0])
	}
	for simbol, words := range at.Expressions {
		if simbol.Value == "A" && words[0].Simbols[1].Value == "f" {
			t.Error("Should've changed 'f' terminal")
		}
	}
}

func TestComplexReplace(t *testing.T) {
	at := models.NewAutomato([]string{
		"S -> ASA | ab | a",
		"A -> b | S",
	})
	fmt.Println("[Start]")

	at = RealocateTerminals(at)
	if len(at.Expressions) != 4 {
		t.Error("Should've generated four expressions")
		at.Explain()
	}
	fmt.Println("[End]")
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
