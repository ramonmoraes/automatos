package models

// GetVariableList should return the list of every possible variables on automato
func (a *Automato) GetVariableList() []Simbol {
	varList := []Simbol{}
	for s, words := range a.Expressions {
		for _, word := range words {
			varList = append(varList, word.GetVariables()...)
		}
		varList = append(varList, s)
	}
	return varList
}
