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
	at := models.NewAutomato([]string{"A -> a|aC", "C -> c | 0"})
	atModified := RemoveEmptyProductions(at)
	if len(atModified.Expressions) != 1 {
		t.Error("Automato should have appended B content on A variable")
	}

	for _, words := range atModified.Expressions {
		if len(words) != 2 {
			t.Error("Automato should have generated two words")
		}
		if words[0].ToString() != "a" || words[1].ToString() != "ac" {
			t.Error("Automato should have words 'a' and 'c'")
		}
	}
	atModified.Explain()
}
