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

// ToString should return a the word's string value
func (w *Word) ToString() string {
	arr := []string{}
	for _, symbol := range w.Simbols {
		arr = append(arr, symbol.Value)
	}
	return strings.Join(arr, "")
}
