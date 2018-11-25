package models

import (
	"fmt"
	"log"
	"strings"
)

// Automato should be a list of expressions
// It can be created with NewAutomato method or normally
type Automato struct {
	Expressions map[Simbol][]Word
}

// NewAutomato should create an Automato parsing the string equations given
func NewAutomato(expressionStringList []string) Automato {
	expressionMap := make(map[Simbol][]Word)
	for _, expressionString := range expressionStringList {
		simbol, word := newExpression(expressionString)
		expressionMap[simbol] = word
	}

	return Automato{
		Expressions: expressionMap,
	}
}

func newExpression(expression string) (Simbol, []Word) {
	splited := strings.Split(expression, "->")
	creator, err := NewSimbol(strings.TrimSpace(splited[0]))

	if err != nil {
		log.Fatal(err)
	}

	generatedList := strings.Split(strings.TrimSpace(splited[1]), "|")
	wordList := []Word{}

	for _, word := range generatedList {
		wordList = append(wordList, NewWord(word))
	}

	return creator, wordList
}

// Explain should describe each automato's expression
func (a *Automato) Explain() {
	fmt.Println("Explaining automato:")
	for simbol, words := range a.Expressions {
		explainedWords := []string{}
		for _, word := range words {
			explainedWords = append(explainedWords, word.ToString())
		}
		fmt.Printf("%s -> %s\n", simbol.Value, strings.Join(explainedWords, " | "))
	}
}

// Fix should correct minor bugs, like
// Trim 'empty simbols' from words with other simbols
// eg: A -> a0, should be A -> a
// and remove duplicated words in same expression
func Fix(a Automato) Automato {
	emptyTrimmedAt := trimEmpty(a)
	singularAt := removeDuplicate(emptyTrimmedAt)
	return singularAt
}

func removeDuplicate(a Automato) Automato {
	newAt := a
	for simbol, words := range a.Expressions {
		uniqueMap := make(map[string]string)
		for _, word := range words {
			str := word.ToString()
			uniqueMap[str] = str
		}

		uniqueWords := []Word{}
		for _, value := range uniqueMap {
			uniqueWords = append(uniqueWords, NewWord(value))
		}
		newAt.Expressions[simbol] = uniqueWords
	}
	return newAt
}

func trimEmpty(a Automato) Automato {
	newAt := a
	for simbol, words := range a.Expressions {
		wordList := []Word{}
		for _, word := range words {
			newStr := word.ToString()
			if len(word.Simbols) > 1 {
				splitedStr := strings.Split(newStr, "0")
				newStr = strings.Join(splitedStr, "")
			}
			wordList = append(wordList, NewWord(newStr))
		}
		newAt.Expressions[simbol] = wordList
	}
	return newAt
}
