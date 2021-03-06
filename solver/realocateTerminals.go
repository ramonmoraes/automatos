package solver

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"../models"
)

// RealocateTerminals should generate Variables for Terminals
func RealocateTerminals(a models.Automato) models.Automato {
	for _, words := range a.Expressions {
		for _, word := range words {
			if len(word.Simbols) > 1 && word.ContainTerminal() {
				return ReplaceFirstTerminalForNewSimbol(a, word)
			}
		}
	}
	return a
}

// ReplaceFirstTerminalForNewSimbol should replace every simbol a automato for another simbol
func ReplaceFirstTerminalForNewSimbol(a models.Automato, wordContainingTerminal models.Word) models.Automato {
	var firstTerminal models.Simbol
	for _, simbol := range wordContainingTerminal.Simbols {
		if !simbol.IsVariable {
			firstTerminal = simbol
			break
		}
	}

	newSimbol := GenerateNewSimbol(a.GetVariableList(), firstTerminal)
	newSimbol = checkingExistingSimbol(a, newSimbol, firstTerminal)

	newStringWord := strings.Replace(wordContainingTerminal.ToString(), firstTerminal.Value, newSimbol.Value, 1)
	replacedAt := replaceStrings(a, wordContainingTerminal.ToString(), newStringWord)
	replacedAt.Expressions[newSimbol] = []models.Word{
		models.Word{
			Simbols: []models.Simbol{firstTerminal},
		},
	}
	return RealocateTerminals(replacedAt)
}

func replaceSimbols(a models.Automato, toBeReplacedSimbol models.Simbol, newSimbol models.Simbol) models.Automato {
	replacedAt := models.Automato{
		Expressions: make(map[models.Simbol][]models.Word),
	}

	for simbol, words := range a.Expressions {
		for _, word := range words {
			wordString := word.ToString()
			replacedString := strings.Replace(wordString, toBeReplacedSimbol.Value, newSimbol.Value, -1)
			newWord := models.NewWord(replacedString)
			_, ok := replacedAt.Expressions[simbol]
			if ok {
				replacedAt.Expressions[simbol] = append(replacedAt.Expressions[simbol], newWord)
			} else {
				replacedAt.Expressions[simbol] = []models.Word{newWord}
			}
		}
	}

	return replacedAt
}

func replaceStrings(a models.Automato, toBeReplacedString string, newString string) models.Automato {
	replacedAt := models.Automato{
		Expressions: make(map[models.Simbol][]models.Word),
	}

	for simbol, words := range a.Expressions {
		for _, word := range words {
			wordString := word.ToString()
			replacedString := strings.Replace(wordString, toBeReplacedString, newString, -1)
			newWord := models.NewWord(replacedString)
			_, ok := replacedAt.Expressions[simbol]
			if ok {
				replacedAt.Expressions[simbol] = append(replacedAt.Expressions[simbol], newWord)
			} else {
				replacedAt.Expressions[simbol] = []models.Word{newWord}
			}
		}
	}

	return replacedAt
}

// GenerateNewSimbol should return a non existing simbol, given the list of already existing simbols
func GenerateNewSimbol(existingVariables []models.Simbol, originatedTerminal models.Simbol) models.Simbol {
	idealSimbol := strings.ToUpper(originatedTerminal.Value)
	varMap := make(map[string]string)
	for _, v := range existingVariables {
		varMap[v.Value] = v.Value
	}

	_, alreadyExist := varMap[idealSimbol]
	if alreadyExist {
		return getRandomSimbol(varMap)
	}
	newSimbol, err := models.NewSimbol(idealSimbol)
	if err != nil {
		log.Fatal(err)
	}
	return newSimbol
}

func getRandomSimbol(varMap map[string]string) models.Simbol {
	rand.NewSource(time.Now().UnixNano())
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXY"
	alpabetSize := len([]rune(alphabet))

	randomPos := rand.Intn(alpabetSize)
	newSimbol, err := models.NewSimbol(strings.Split(alphabet, "")[randomPos])
	if err != nil {
		log.Fatal(err)
	}
	return newSimbol
}

func checkingExistingSimbol(a models.Automato, newSimbol models.Simbol, value models.Simbol) models.Simbol {
	for simbol, words := range a.Expressions {
		if len(words) == 1 && len(words[0].Simbols) == 1 && words[0].Simbols[0] == value {
			fmt.Printf("Found :%s generating %s already\n", simbol.Value, value.Value)
			return simbol
		}
	}
	return newSimbol
}
