package models

import (
	"testing"
)

func TestCreateExpression(t *testing.T) {
	ex, err := NewExpression("A -> ab | a")
	if err != nil {
		t.Error("NewExpression should not have err")
	}

	creator := ex.Variable
	generated := ex.Words
	print(generated)

	if creator.Value != "A" {
		t.Error("NewExpression should have creater a creator with value of A")
	}

	if len(generated) != 2 {
		t.Error("NewExpression have generated with len of 1")
	}
}
