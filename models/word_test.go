package models

import "testing"

func TestVariablesFunction(t *testing.T) {
	w := NewWord("aB")
	if w.GetTerminals()[0].Value != "a" {
		t.Error("Expected 'w' to have the a terminal")
	}

	w = NewWord("A")
	if len(w.GetTerminals()) != 0 {
		t.Error("Expected 'w' to not have terminal")
	}
}

func TestVariableFunction(t *testing.T) {
	w := NewWord("aB")
	if w.GetVariables()[0].Value != "B" {
		t.Error("Expected 'w' to have the a Variable")
	}

	w = NewWord("b")
	if len(w.GetVariables()) != 0 {
		t.Error("Expected 'w' to not have Variable")
	}
}

func TestIsEmpty(t *testing.T) {
	w := NewWord("a0")
	if w.IsEmpty() {
		t.Error("Word a0 should not be empty")
	}

	w = NewWord("0")
	if !w.IsEmpty() {
		t.Error("Word 0 should be empty")
	}
}
