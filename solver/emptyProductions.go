package solver

import (
	"fmt"
	"strconv"

	"../models"
)

// RemoveEmptyProductions should remove every expression that generates empty word
func RemoveEmptyProductions(automato models.Automato) models.Automato {
	atWithoutProductions := removeEmptyProductions(automato)
	fixedAt := models.Fix(atWithoutProductions)
	return fixedAt
}

func removeEmptyProductions(automato models.Automato) models.Automato {
	for simbol, words := range automato.Expressions {
		for _, word := range words {
			if word.IsEmpty() {
				return replaceEmptySimbolCombination(automato, simbol)
			}
		}
	}

	return automato
}

func replaceEmptySimbolCombination(automato models.Automato, simbolWithEmptyWord models.Simbol) models.Automato {
	at := automato
	for simbolIndex, words := range automato.Expressions {
		var newWords []models.Word
		for _, word := range words {
			for _, wordSimbol := range word.Simbols {
				if wordSimbol == simbolWithEmptyWord {
					combinations := getCombination(word, simbolWithEmptyWord)
					newWords = append(newWords, combinations...)
				} else {
					newWords = append(newWords, word)
				}
			}
		}
		at.Expressions[simbolIndex] = newWords
	}
	at.Expressions[simbolWithEmptyWord] = removeEmptyFrom(at.Expressions[simbolWithEmptyWord])
	return removeEmptyProductions(at)
}

func getCombination(word models.Word, simbol models.Simbol) []models.Word {
	var equalValueIndices []int
	for indice, simb := range word.Simbols {
		if simb.Value == simbol.Value {
			equalValueIndices = append(equalValueIndices, indice)
		}
	}

	var combinations []models.Word
	binaryCombination := getBitSlice(len(equalValueIndices))
	for _, bitNumber := range binaryCombination {
		nw := getCombinationBasedOnBinary(word, simbol, equalValueIndices, bitNumber)
		combinations = append(combinations, nw)
	}
	return combinations
}

func getCombinationBasedOnBinary(word models.Word, simbol models.Simbol, equalValueIndices []int, binaryNumber string) models.Word {
	nw := models.NewWord(word.ToString()) // Could not find a better way to not mutate word
	for bitIndice, bit := range binaryNumber {
		sx := equalValueIndices[bitIndice]
		if string(bit) == "0" {
			nw.Simbols[sx].Value = "0"
		}
	}
	return nw
}

func getBitSlice(number int) []string {
	var bitSlice []string
	i := 0
	highestBitNumber := strconv.FormatInt(int64(number*number), 2)
	padSize := len(highestBitNumber) - 1
	for i < number*number {
		binValue := strconv.FormatInt(int64(i), 2)
		bitString := fmt.Sprintf("%0*s", padSize, binValue)
		bitSlice = append(bitSlice, bitString)
		i = i + 1
	}
	return bitSlice
}

func removeEmptyFrom(words []models.Word) []models.Word {
	var nw []models.Word
	for _, w := range words {
		if !w.IsEmpty() {
			nw = append(nw, w)
		}
	}
	return nw
}

func explain(words []models.Word) {
	for _, w := range words {
		fmt.Println(w.ToString())
	}
}
