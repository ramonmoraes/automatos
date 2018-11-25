package solver

import (
	"testing"

	"../models"
)

func TestUselessSimbols(t *testing.T) {
	at := models.NewAutomato([]string{"A -> B | c", "B -> cd | f | B"})
	atModified := UselessSimbol(at)
	for simbol, words := range atModified.Expressions {
		if simbol.Value == "A" &&
			words[0].ToString() != "cd" &&
			words[1].ToString() != "f" &&
			words[2].ToString() != "c" {
			t.Error("A is generating incorrectly")
		}
		if simbol.Value == "B" &&
			words[0].ToString() != "cd" &&
			words[2].ToString() != "f" {
			t.Error("B is generating incorrectly")
		}

	}
}

func TestRemoveSingleRecursive(t *testing.T) {
	at := models.NewAutomato([]string{"A -> as | bfg | A", "B -> b | xz | B"})
	at = removeSingleRecursive(at)
	for simbol, words := range at.Expressions {
		for _, word := range words {
			for _, simb := range word.Simbols {
				if len(word.Simbols) == 1 && simb == simbol {
					t.Error("Generator simbol shold not exist in itself")
				}
			}
		}
	}
}
