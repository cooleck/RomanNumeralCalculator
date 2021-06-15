package structs

const (
	IsNum            = 1 << iota
	IsBinaryOperator = 1 << iota
	IsUnaryOperator  = 1 << iota
	IsOpenBracket    = 1 << iota
	IsCloseBracket   = 1 << iota
)

type Token struct {
	Precedence int
	Val        interface{}
	Type       int
}

func NewToken(precedence int, val interface{}, t int) *Token {
	return &Token{Precedence: precedence, Val: val, Type: t}
}
