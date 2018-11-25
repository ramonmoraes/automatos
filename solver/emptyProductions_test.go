package solver

import (
	"fmt"
	"testing"

	"../models"
)

func TestRemoveSimpleEmptyProduction(t *testing.T) {
	at := models.NewAutomato([]string{"A -> a | BaB", "B -> 0 | b"})
	atModified := RemoveEmptyProductions(at)
	if len(atModified.Expressions) != 2 {
		t.Error("Automato should have two expressions")
	}

	for simb, words := range atModified.Expressions {
		if simb.Value == "A" && len(words) != 4 {
			t.Error("A -> Should have generated 4 words")
		}

		if simb.Value == "B" && len(words) != 1 {
			t.Error("B -> Should have generated 1 word")
		}
	}
}

func TestCombination(t *testing.T) {
	word := models.NewWord("DCD")
	simb, _ := models.NewSimbol("D")

	comb := getCombination(word, simb)
	if len(comb) != 4 {
		t.Error("Should have generated 4 combinations")
	}

	if comb[0].ToString() != "0C0" &&
		comb[0].ToString() != "0CD" &&
		comb[0].ToString() != "DC0" &&
		comb[0].ToString() != "DCD" {
		t.Error("Should generate correct combination")
		explain(comb)
	}
}

func TestRemoveEmptyFrom(t *testing.T) {
	at := models.NewAutomato([]string{
		"A -> B | 0",
		"B -> b | 0",
	})

	for _, words := range at.Expressions {
		nw := removeEmptyFrom(words)
		for _, w := range nw {
			for _, simbol := range w.Simbols {
				if simbol.IsEmpty {
					t.Error("Should not have any empty")
					fmt.Println(w.ToString())
				}
			}
		}
	}
}
