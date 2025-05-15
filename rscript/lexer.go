package rscript

import (
	"strings"
)

func Tokenize(script string) ([]Action, error) {
	lines := strings.Split(script, "\n")
	tokens := []Action{}
	for i := 0; i < len(lines); i++ {
		words := strings.Split(lines[i], " ")

		switch words[0] {
		case "//":
			continue
		case "let":
			varName := words[1]
			// var expr Expression
			if words[2] == "=" {
				//expr := ExpressionBoolOperator{Expr: ""}
			}
			tokens = append(tokens, ActionNewVar{
				VarName: varName,
				//AssignmentExpression: ,
			})
		}

	}
	return tokens, nil
}

func (action ActionNewVar) Run(rt *Runtime) error {
	rt.setVar(Variable{action.VarName, action.Run(rt)})
	return nil
}
