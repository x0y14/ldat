package python

import "testing"

func TestTokenizer_Tokenize(t *testing.T) {
	pyscript := `print("hello python")`

	ast, err := DumpAstWithCode(pyscript)
	if err != nil {
		t.Fatal(err)
	}

	tokenizer := NewTokenizer()
	err = tokenizer.Tokenize(ast)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTokenizer_Tokenize2(t *testing.T) {
	pyscriptPath := `
class Name:
    def __init__(self, first: str, middle: str, last: str):
        self.first: str = first
        self.middle: str = middle
        self.last: str = last

    def to_s(self) -> None:
        print(self.first + " " + self.middle + " " + self.last)
`
	ast, err := DumpAstWithCode(pyscriptPath)
	if err != nil {
		t.Fatal(err)
	}

	tokenizer := NewTokenizer()
	err = tokenizer.Tokenize(ast)
	if err != nil {
		t.Fatal(err)
	}
}
