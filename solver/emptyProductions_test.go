package solver

import (
	"testing"

	"../models"
)

func TestRemoveSimpleEmptyProduction(t *testing.T) {
	at := models.NewAutomato([]string{"A -> a", "B -> 0"})
	atModified := RemoveEmptyProductions(at)
	if len(atModified.Expressions) != 1 {
		t.Error("Automato should have only one expression")
		atModified.Explain()
	}
}

func TestRemoveDoubleEmptyProduction(t *testing.T) {
	at := models.NewAutomato([]string{"A -> a|B", "B -> c | 0"})
	atModified := RemoveEmptyProductions(at)
	if len(atModified.Expressions) != 2 {
		t.Error("Automato should have only one expression")
		atModified.Explain()
	}
}
