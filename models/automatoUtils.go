package models

// GetVariableList should return the list of every possible variables on automato
// func (a *Automato) GetVariableList() []Simbol {
// 	varList := []Simbol{}
// 	for simbol, _ := range a.Expressions {
// 		varList = append(varList, expression.Variable)
// 	}
// 	return varList
// }

// // GetPossibleCreatedsSimbols should return the list of every simbol that can be created
// func (a *Automato) GetPossibleCreatedsSimbols() []Simbol {
// 	wordList := []Word{}
// 	for _, expression := range a.Expressions {
// 		wordList = append(wordList, expression.Words...)
// 	}

// 	simbolList := []Simbol{}
// 	for _, word := range wordList {
// 		simbolList = append(simbolList, word.Simbols...)
// 	}
// 	return simbolList
// }

// // GetEverySimbol should return the list of every possible variables on automato
// func (a *Automato) GetEverySimbol() []Simbol {
// 	return append(a.GetVariableList(), a.GetPossibleCreatedsSimbols()...)
// }

// // Explain should call explain for every expression in automato
// func (a *Automato) Explain() {
// 	fmt.Println("Explaining Automato:")
// 	for _, expression := range a.Expressions {
// 		expression.Explain()
// 	}
// }
