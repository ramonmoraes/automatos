package models

import (
	"testing"
)

func testCreateExpression(t *testing.T) {
	ex, err := NewExpression("A -> ab")
	if err != nil {
		t.Error("NewExpression should not have err")
	}

	creator := ex.Creator
	generated := ex.Generated
	print(generated)

	if creator.Value != "A" {
		t.Error("NewExpression should have creater a creator with value of A")
	}

	if len(generated) != 2 {
		t.Error("NewExpression have generated with len of 2")
	}
}
