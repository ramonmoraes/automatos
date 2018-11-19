package solver

// import (
// 	"../models"
// )

// // UselessSimbol should remove every useless simbol
// func UselessSimbol(automato models.Automato) models.Automato {
// 	validExpressionList := []models.Expression{}
// 	for _, expr := range automato.Expressions {
// 		for _, word := range expr.Words {
// 			if word.ContainTerminal() {
// 				validExpressionList = append(validExpressionList, expr)
// 			}
// 		}
// 	}

// 	return models.Automato{
// 		Expressions: validExpressionList,
// 	}
// }
