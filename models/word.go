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

// ContainVariable should return a list of variables on the word
func (w *Word) ContainVariable() bool {
	return len(w.GetVariables()) != 0
}

// GetVariables should return a list of variables on the word
func (w *Word) GetVariables() []Simbol {
	varList := []Simbol{}
	for _, simb := range w.Simbols {
		if simb.IsVariable {
			varList = append(varList, simb)
		}
	}
	return []Simbol{}
}

// ContainTerminal should return a list of non-variables on the word
func (w *Word) ContainTerminal() bool {
	return len(w.GetTerminals()) != 0
}

// GetTerminals should return a list of non-variables on the word
func (w *Word) GetTerminals() []Simbol {
	terminalList := []Simbol{}
	for _, simb := range w.Simbols {
		if !simb.IsVariable {
			terminalList = append(terminalList, simb)
		}
	}
	return []Simbol{}
}

// ToString should return a the word's string value
func (w *Word) ToString() string {
	arr := []string{}
	for _, symbol := range w.Simbols {
		arr = append(arr, symbol.Value)
	}
	return strings.Join(arr, "")
}
