package python

import (
	"fmt"
	"regexp"
)

type Tokenizer struct {
	pos   int
	runes []rune
}

/*

age: int = 19
name: str = "john"
print("hello" + name + str(age))

Module(
    body=[
        AnnAssign(target=Name(id='age', ctx=Store()), annotation=Name(id='int', ctx=Load()),value=Constant(value=19, kind=None), simple=1),
        AnnAssign(target=Name(id='name', ctx=Store()), annotation=Name(id='str', ctx=Load()),value=Constant(value='john', kind=None), simple=1),
        Expr(
            value=Call(
                func=Name(id='print', ctx=Load()),
                args=[
                    BinOp(
                        left=BinOp(left=Constant(value='hello, ', kind=None), op=Add(), right=Name(id='name', ctx=Load())),
                        op=Add(),
                        right=Call(func=Name(id='str', ctx=Load()), args=[Name(id='age', ctx=Load())], keywords=[])
                    )
                ],
                keywords=[])
        )
    ],
    type_ignores=[]
)

^Module\([\s]+body\=\[(.*)\]\,[\s]+type_ignores\=\[(.*)\][\s]+\)$

$0 body=[(.*)]
$1 type_ignores=[(.*)]

*/

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		pos:   0,
		runes: nil,
	}
}

func (t *Tokenizer) isEof() bool {
	return t.pos >= len(t.runes)
}

func (t *Tokenizer) curt() rune {
	return t.runes[t.pos]
}

func (t *Tokenizer) goNext() {
	t.pos++
}

func (t *Tokenizer) load(code string) error {

	reg := regexp.MustCompile(`Module\(\s*body=\[(.*)],\s*type_ignores=\[(.*)]\s*\)\s*`)

	result := reg.FindStringSubmatch(code)
	if len(result) != 3 {
		return fmt.Errorf("ast format error")
	}

	t.runes = []rune(result[1])
	return nil
}

func (t *Tokenizer) Tokenize(code string) (*Token, error) {
	err := t.load(code)
	if err != nil {
		return nil, err
	}

	var result *Token = &Token{SOF, "", 0, 0, nil}

	for !t.isEof() {
		//fmt.Printf("%v", string(t.curt()))
		//t.goNext()
		switch {
		case 65 <= t.curt() && t.curt() <= 90 || 97 <= t.curt() && t.curt() <= 122:
			// A-Z || a-z
			t.ident()
			//case t.curt() == '(' || t.curt() == ')':
			//	t.goNext()
			//case t.curt() == '[' || t.curt() == ']':
			//	t.goNext()
			//case t.curt() == '=':
			//	t.goNext()
		}
	}

	return result, nil
}

func (t *Tokenizer) ident() {}
