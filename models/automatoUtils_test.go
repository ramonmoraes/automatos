package models

import (
	"fmt"
	"testing"
)

func TestGetEveryVariable(t *testing.T) {
	at := NewAutomato([]string{"A -> eB", "B -> dF", "F -> A | 0"})
	varList := at.GetVariableList()
	if len(varList) != 3 {
		t.Error("Should have 3 variables")
		for _, v := range varList {
			fmt.Println(v.Value)
		}
	}
}
