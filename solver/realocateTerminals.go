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
	varList := a.GetVariableList()
	for _, words := range a.Expressions {
		for _, word := range words {
			if len(word.Simbols) > 1 && word.ContainTerminal() {
				for _, wordSimbol := range word.Simbols {
					if !wordSimbol.IsVariable {
						newSimbol := GenerateNewSimbol(varList, wordSimbol)
						fmt.Println(newSimbol)
					}
				}
			}
		}
	}
	return a
}

// GenerateNewSimbol should return a non existing simbol, given the list of already existing simbols
func GenerateNewSimbol(existingVarialbes []models.Simbol, originatedTerminal models.Simbol) models.Simbol {
	idealSimbol := strings.ToUpper(originatedTerminal.Value)
	varMap := make(map[string]string)
	for _, v := range existingVarialbes {
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
	fmt.Println(randomPos)

	newSimbol, err := models.NewSimbol(strings.Split(alphabet, "")[randomPos])
	if err != nil {
		log.Fatal(err)
	}
	return newSimbol
}
