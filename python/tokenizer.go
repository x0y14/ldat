package python

import "fmt"

type Tokenizer struct {
	pos   int
	runes []rune
}

/*

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

func (t *Tokenizer) load(code string) {
	t.runes = []rune(code)
}

func (t *Tokenizer) Tokenize(code string) {
	t.load(code)

	for !t.isEof() {
		fmt.Printf("%v", string(t.curt()))
		t.goNext()
	}
}
