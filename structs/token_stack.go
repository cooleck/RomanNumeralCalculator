package structs

import "errors"

var (
	ErrBadToken          = errors.New("the token is bad")
	ErrUnmatchedBrackets = errors.New("there are unmatched brackets in the expression")
)

type TokenStack []*Token

func NewTokenStack(n int) TokenStack {
	return make(TokenStack, 0, n)
}

func (s *TokenStack) Top() *Token {

	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *TokenStack) Pop() {
	if len(*s) == 0 {
		return
	}

	*s = (*s)[:len(*s)-1]
}

func (s *TokenStack) Empty() bool {
	return len(*s) == 0
}

func (s *TokenStack) Push(t *Token, RPN *[]*Token) (err error) {
	switch t.Type {
	case IsOpenBracket:
		*s = append(*s, t)
		return nil
	case IsBinaryOperator, IsUnaryOperator:
		err = s.pushOperator(t, RPN)
		return err
	default:
		return ErrBadToken
	}
}

func (s *TokenStack) pushOperator(t *Token, RPN *[]*Token) (err error) {
	for top := s.Top(); top != nil && top.Type&(IsUnaryOperator|IsBinaryOperator) != 0 && top.Precedence >= t.Precedence; top = s.Top() {
		*RPN = append(*RPN, top)
		s.Pop()
	}
	*s = append(*s, t)
	return nil
}

func (s *TokenStack) Drop(RPN *[]*Token) (err error) {
	for !s.Empty() {
		top := s.Top()
		if top.Type&(IsUnaryOperator|IsBinaryOperator) == 0 {
			return ErrUnmatchedBrackets
		}
		*RPN = append(*RPN, top)
		s.Pop()
	}
	return nil
}
