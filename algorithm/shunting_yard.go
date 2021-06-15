package algorithm

import (
	"errors"
	"vk/structs"
)

var (
	ErrUnidentifiedTokenType = errors.New("unidentified token type")
)

func ShuntingYard(infixNotation []*structs.Token) (RPN []*structs.Token, err error) {
	RPN = make([]*structs.Token, 0, len(infixNotation))
	stack := structs.NewTokenStack(0)

	for _, token := range infixNotation {
		switch token.Type {
		case structs.IsNum:
			RPN = append(RPN, token)
		case structs.IsBinaryOperator, structs.IsUnaryOperator, structs.IsOpenBracket:
			err = stack.Push(token, &RPN)
			if err != nil {
				return nil, err
			}
		case structs.IsCloseBracket:
			for top := stack.Top(); top != nil && top.Type != structs.IsOpenBracket; top = stack.Top() {
				RPN = append(RPN, top)
				stack.Pop()
			}

			if stack.Top() == nil {
				return nil, structs.ErrUnmatchedBrackets
			}
			stack.Pop()
		default:
			return nil, ErrUnidentifiedTokenType
		}
	}

	err = stack.Drop(&RPN)
	return RPN, err
}
