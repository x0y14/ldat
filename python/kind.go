package python

type TokenKind int

const (
	_ TokenKind = iota
	SOF
	EOF
	IDENT  // keyword
	STRING // ""
	WHITE  // " "
	LBR    // (
	RBR    // )
	LSQB   // [
	RSQB   // ]
	COMMA  // ,
	EQU    // =
)

var kinds = [...]string{
	IDENT:  "IDENT",
	STRING: "STRING",
	WHITE:  "WHITE",
	LBR:    "LBR",
	RBR:    "RBR",
	LSQB:   "LSQB",
	RSQB:   "RSQB",
	COMMA:  "COMMA",
	EQU:    "EQU",
}

func (k TokenKind) String() string {
	return kinds[k]
}
