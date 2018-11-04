package models

import (
	"log"
	"strings"
)

// Word should only be a list of simbols
type Word struct {
	Simbols []Simbol
}

// NewWord should create a WordType baed on given string
func NewWord(word string) Word {
	trimmedWord := strings.TrimSpace(word)
	letters := strings.Split(trimmedWord, "")
	newWord := Word{}
	for _, letter := range letters {
		simbol, err := NewSimbol(letter)
		if err != nil {
			log.Fatal(err)
		}
		newWord.Simbols = append(newWord.Simbols, simbol)
	}
	return newWord
}

// ContainsVariable should return a list of variables on the word
func (w *Word) ContainsVariable() ([]Simbol, bool) {
	varList := []Simbol{}
	for _, simb := range w.Simbols {
		if simb.IsVariable {
			varList = append(varList, simb)
		}
	}
	return []Simbol{}, len(varList) == 0
}

// ContainsTerminals should return a list of non-variables on the word
func (w *Word) ContainsTerminals() ([]Simbol, bool) {
	terminalList := []Simbol{}
	for _, simb := range w.Simbols {
		if !simb.IsVariable {
			terminalList = append(terminalList, simb)
		}
	}
	return []Simbol{}, len(terminalList) == 0
}

// ToString should return a the word's string value
func (w *Word) ToString() string {
	arr := []string{}
	for _, symbol := range w.Simbols {
		arr = append(arr, symbol.Value)
	}
	return strings.Join(arr, "")
}
