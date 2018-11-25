package solver

import (
	"testing"

	"../models"
)

func TestRemoveSimpleEmptyProduction(t *testing.T) {
	at := models.NewAutomato([]string{"A -> a | B", "B -> 0"})
	atModified := RemoveEmptyProductions(at)
	if len(atModified.Expressions) != 1 {
		t.Error("Automato should have only one expression")
	}
	atModified.Explain()
}

func TestRemoveDoubleEmptyProduction(t *testing.T) {
	at := models.NewAutomato([]string{"A -> a|B", "B -> c | 0"})
	atModified := RemoveEmptyProductions(at)

	if len(atModified.Expressions) != 1 {
		t.Error("Automato should have appended B content on A variable")
	}

	for _, words := range atModified.Expressions {
		if len(words) != 2 {
			t.Error("Automato should have generated two words")
		}
		if words[0].ToString() != "a" || words[1].ToString() != "c" {
			t.Error("Automato should have words 'a' and 'c'")
		}
	}

	atModified.Explain()
}

func TestRemoveComplex(t *testing.T) {
	at := models.NewAutomato([]string{
		"A -> B | aC",
		"B -> d | C",
		"C -> c | 0",
	})
	atModified := RemoveEmptyProductions(at)
	if len(atModified.Expressions) != 2 {
		t.Error("Automato should have two expressions")
	}

	for Simbol, words := range atModified.Expressions {
		if Simbol.Value == "A" {
			if words[0].ToString() != "a" &&
				words[1].ToString() != "B" &&
				words[1].ToString() != "ac" {
				t.Error("Automato should have words 'a' and 'B' and 'ac'")
				atModified.Explain()
			}
		}
		if Simbol.Value == "B" {
			if words[0].ToString() != "d" &&
				words[1].ToString() != "c" {
				t.Error("Automato should have words 'c' and 'd'")
				atModified.Explain()
			}
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
