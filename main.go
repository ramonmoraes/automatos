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
		var atExpressions []string
		var inputText string
		fmt.Println("Write expressions:")

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
		chomAt := solver.Solve(at, *verbose)
		chomAt.Explain()
	}
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
