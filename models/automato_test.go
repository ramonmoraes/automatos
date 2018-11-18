package models

import (
	"testing"
)

func TestCreateExpression(t *testing.T) {
	simbol, words := newExpression("A -> ab | a")
	if simbol.Value != "A" {
		t.Error("expression should have A as variable")
	}

	if len(words) != 2 {
		t.Error("newExpression should have generated two words")
	}

	if words[0].ToString() != "ab" && words[1].ToString() != "a" {
		t.Error("newExpression have created correct words")
	}

}
