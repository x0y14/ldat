package python

import "testing"

func TestTokenizer_Tokenize(t *testing.T) {
	pyscript := `print("hello python")`

	tokenizer := NewTokenizer()
	tokenizer.Tokenize(pyscript)
}
