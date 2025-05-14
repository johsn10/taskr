package rscript

type Runtime struct {
	Script string
}

func NewRuntime(script string) Runtime {
	return Runtime{Script: script}
}
func (rt Runtime) Run() {

}
