package ldat

type IRCodeGen struct {
	Ops []*CommonOperation
}

func NewIRCodeGen() *IRCodeGen {
	return &IRCodeGen{
		Ops: nil,
	}
}
