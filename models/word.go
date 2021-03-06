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
	return varList
}

// ContainTerminal should return a list of non-variables on the word
func (w *Word) ContainTerminal() bool {
	return len(w.GetTerminals()) != 0
}

// GetTerminals should return a list of non-variables on the word
func (w *Word) GetTerminals() []Simbol {
	terminalList := []Simbol{}
	for _, simb := range w.Simbols {
		if !simb.IsVariable && !simb.IsEmpty {
			terminalList = append(terminalList, simb)
		}
	}
	return terminalList
}

// ToString should return a the word's string value
func (w *Word) ToString() string {
	arr := []string{}
	for _, symbol := range w.Simbols {
		arr = append(arr, symbol.Value)
	}
	return strings.Join(arr, "")
}

// IsEmpty should return if the word contains only one empty simbol
func (w *Word) IsEmpty() bool {
	return !w.ContainVariable() && !w.ContainTerminal()
}

// Replace should return a new word replacing the original simbol:first param
// for the new simbols:second param
func (w *Word) Replace(originalSimbol Simbol, newSimbols []Simbol) Word {
	var newWord string
	for _, simbol := range newSimbols {
		newWord += simbol.Value
	}

	newString := strings.Replace(
		w.ToString(),
		originalSimbol.Value,
		newWord,
		-1,
	)

	return NewWord(newString)
}

// ReplaceWords should return words replacing the original simbol:first param
// for the new simbols:second param
// Entrada:
// A -> a | aB
// B -> D | E
// Saida:
// A -> a | aD | AE
func (w *Word) ReplaceWords(originalSimbol Simbol, words []Word) []Word {
	wordList := []Word{}
	for _, word := range words {
		refWord := w.Replace(originalSimbol, word.Simbols)
		wordList = append(wordList, refWord)
	}
	return wordList
}

// Contains should return if the word contains the simbol
func (w *Word) Contains(simbol Simbol) bool {
	for _, s := range w.Simbols {
		if s == simbol {
			return true
		}
	}
	return false
}
