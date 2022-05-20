package python

type Token struct {
	TokenKind
	Raw  string
	S    int
	E    int
	Next *Token
}
