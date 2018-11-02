package parser

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

	if creator.Value != "A" {
		t.Error("NewExpression should have creater a creator with value of A")
	}

	if len(generated) != 2 {
		t.Error("NewExpression have generated with len of 2")
	}
	if generated[0].Value != "a" {
		t.Error("NewExpression should have generated simbol with value of a")
	}

	if generated[1].Value != "b" {
		t.Error("NewExpression should have generated simbol with value of b")
	}
}
