package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//Maniupulating Values
	COPY = "COPY"
	ADDI = "ADDI"
	SUBI = "SUBI"
	MULI = "MULI"
	DIVI = "DIVI"
	MODI = "MODI"
	SWIZ = "SWIZ"

	//Branching
	MARK = "MARK"
	JUMP = "JUMP"
	TJMP = "TJMP"
	FJMP = "FJMP"

	//Testing Values
	TEST = "TEST"

	//Lifecycle
	REPL = "REPL"
	HALT = "HALT"
	KILL = "KILL"

	//Movement
	HOST = "HOST"
	LINK = "LINK"

	//Communication
	MODE = "MODE"
	VOID = "VOID"

	//File Manipulation
	MAKE = "MAKE"
	GRAB = "GRAB"
	FILE = "FILE"
	SEEK = "SEEK"
	DROP = "DROP"
	WIPE = "WIPE"

	//MISC
	NOTE = "NOTE"
	NOOP = "NOOP"
	RAND = "RAND"

	INT = "INT"

	F = "F"
	X = "X"

	LABEL   = "LABEL"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

func LookupLabel(label string) TokenType {
	if tok, ok := keywords[label]; ok {
		return tok
	}
	return LABEL
}

var keywords = map[string]TokenType{
	"MARK": MARK,
	"JUMP": JUMP,
	"TJMP": TJMP,
	"FJMP": FJMP,
}
