package rscript

type Variable struct {
	name  string
	value interface{}
}
type Runtime struct {
	Script    string
	variables []Variable
}

func (rt *Runtime) setVar(variable Variable) {
	rt.variables = append(rt.variables, variable)
}

func NewRuntime(script string) Runtime {
	return Runtime{Script: script}
}
func (rt Runtime) Run() {
	//tokens := Tokenize(rt.Script)
}
