package solver

import (
	"testing"

	"../models"
)

func TestRealocateVariables(t *testing.T) {
	at := models.NewAutomato([]string{
		"A -> BCDE",
	})

	at = RealocateVariables(at)
	if len(at.Expressions) != 3 {
		t.Error("Should have generated 3 expressions")
	}
}
