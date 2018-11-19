package models

import (
	"testing"
)

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

func TestReplace1(t *testing.T) {
	w := NewWord("aB")
	initial, _ := NewSimbol("B")
	simbolList := []Simbol{}
	for _, s := range []string{"d", "F"} {
		simb, _ := NewSimbol(s)
		simbolList = append(simbolList, simb)
	}

	w.Replace(initial, simbolList)
	if len(w.Simbols) != 3 {
		t.Error("Should replace simbols correctly")
	}

	if w.ToString() != "adF" {
		t.Error("Word should be adF")
	}
}

func TestReplace2(t *testing.T) {
	w := NewWord("aBc")
	initial, _ := NewSimbol("B")
	simbolList := []Simbol{}
	for _, s := range []string{"d", "F"} {
		simb, _ := NewSimbol(s)
		simbolList = append(simbolList, simb)
	}

	w.Replace(initial, simbolList)
	if len(w.Simbols) != 4 {
		t.Error("Should replace simbols correctly")
	}

	if w.ToString() != "adFc" {
		t.Error("Word should be adFc")
		t.Log(w.ToString())
	}
}

func TestReplace3(t *testing.T) {
	w := NewWord("aaB")
	initial, _ := NewSimbol("a")
	simbolList := []Simbol{}
	for _, s := range []string{"e", "e"} {
		simb, _ := NewSimbol(s)
		simbolList = append(simbolList, simb)
	}

	w.Replace(initial, simbolList)
	if len(w.Simbols) != 5 {
		t.Error("Should replace simbols correctly")
	}

	if w.ToString() != "eeeeB" {
		t.Error("Word should be eeeeB")
		t.Log(w.ToString())
	}
}

func TestReplace(t *testing.T) {
	w := NewWord("aBc")
	w2 := NewWord("d")
	w3 := NewWord("eF")

	initial, _ := NewSimbol("B")
	nw := w.ReplaceWords(initial, []Word{w2, w3})
	if len(nw) != 2 {
		t.Error("ReplaceWords shoudl return 2 words")
	}

	if nw[0].ToString() != "adc" || nw[1].ToString() != "aeFc" {
		t.Error("ReplaceWords should replace simbols correctly")
	}
}
