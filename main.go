package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"./models"
	"./solver"
)

func main() {
	print("\033[H\033[2J") // clean console
	runExample := flag.Bool("example", false, "Run example func instead of given inputs")
	verbose := flag.Bool("verbose", false, "When true, it outputs every step of the algorithm")
	flag.Parse()

	if *runExample {
		example(*verbose)
	} else {
		input(*verbose)
	}
}

func input(verbose bool) {
	var atExpressions []string
	var inputText string
	fmt.Println("Input disclaimers:")
	fmt.Println("1 - The program will start runing after ocurring a empty input")
	fmt.Printf("2 - The automato should follow this pattern:\n    X -> a | B\n    B -> c\n")
	fmt.Println("3 - The program will not consider S as the 'start' argument, the first one will be considered the start instead")
	fmt.Println("4 - The alphabet have to respect the range between 'a' to 'Z'")
	fmt.Println("5 - Numbers are invalid, except '0' which represents 'empty'")
	fmt.Println("Waiting for input...")
	for inputText != "\n" {
		reader := bufio.NewReader(os.Stdin)
		inputText, _ = reader.ReadString('\n')
		if inputText != "\n" {
			atExpressions = append(atExpressions, inputText)
		}
	}

	fmt.Println("[Initial Automato]")
	at := models.NewAutomato(atExpressions)
	at.Explain()

	fmt.Println("[End Automato]")
	chomAt := solver.Solve(at, verbose)
	chomAt.Explain()
}

func example(verbose bool) {
	at := models.NewAutomato([]string{
		"S -> ASA | aB ",
		"A -> B | S",
		"B -> b | 0",
	})

	fmt.Println("[Initial Automato]")
	at.Explain()

	chomAt := solver.Solve(at, verbose)
	fmt.Println("[End Automato]")
	chomAt.Explain()
}
