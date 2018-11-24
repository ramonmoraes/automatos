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
	at.Explain()
}
