package rscript

type Action interface {
	Run(*Runtime) error
}

type Expression interface {
	Run() (interface{}, error)
}

type ExpressionBash struct {
	Bash string
}

type ExpressionBoolOperator struct {
	Operator string
	Value1   bool
	Value2   bool
}

type ExpressionString struct {
	Expr string
}

// Actions
type ActionBashCommand struct {
	Bash string
}

type ActionNewVar struct {
	VarName              string
	AssignmentExpression Expression
}

type ActionIfStatement struct {
	Expression Expression
	Operator   string
}

type ActionIfBranch struct {
	WasTrue bool
	Action  Action
}
