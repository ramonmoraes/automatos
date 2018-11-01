package parser

import (
	"testing"
)

func testCreateTransformation(t *testing.T) {
	ex, err := NewTrasformation("A -> ab")
	if err != nil {
		t.Error("NewTrasformation should not have err")
	}

	creator := ex.Creator
	generated := ex.Generated

	if creator.Value != "A" {
		t.Error("NewTrasformation should have creater a creator with value of A")
	}

	if len(generated) != 2 {
		t.Error("NewTrasformation have generated with len of 2")
	}
	if generated[0].Value != "a" {
		t.Error("NewTrasformation should have generated simbol with value of a")
	}

	if generated[1].Value != "b" {
		t.Error("NewTrasformation should have generated simbol with value of b")
	}
}
