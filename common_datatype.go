package ldat

type Datatype int

const (
	Unknown Datatype = iota

	Void

	String
	Char

	Float
	Double

	I64
	I32
	Ui64
	Ui32
)
