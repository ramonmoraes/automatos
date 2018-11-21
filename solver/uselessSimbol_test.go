package solver

import (
	"testing"

	"../models"
)

func TestUselessSimbols(t *testing.T) {
	at := models.NewAutomato([]string{"A -> B | c", "B -> C", "C -> D", "D -> e"})
	atModified := UselessSimbol(at)

	if len(atModified.Expressions) != 2 {
		t.Error("Automato should have two expressions")
	}

	for simbol, words := range atModified.Expressions {
		if simbol.Value == "A" {
			if words[0].ToString() != "B" &&
				words[1].ToString() != "c" {
				t.Error("Automato should have words 'B' and 'ac'")
				atModified.Explain()
			}
		}
		if simbol.Value == "B" {
			if words[0].ToString() != "e" {
				t.Error("Automato should have words 'c'")
				atModified.Explain()
			}
		}
	}
}
