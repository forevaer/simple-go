package main

type OP interface {
	do(int) error
}

type OPInstance struct {
	creator func(int) error
}

func (op OPInstance) do(a int) error {
	return op.creator(a)
}
