package algorithm

import (
	"errors"
	"math"
	"vk/structs"
)

var (
	ErrOperandMissed         = errors.New("the operand of operator is missed")
	ErrUnknownBinaryOperator = errors.New("unknown binary operator")
	ErrOperatorMissed        = errors.New("the operator missed")
	ErrDivideByZero          = errors.New("division by zero ")
	ErrInt64Overflow         = errors.New("int64 overflow")
)

func Calculate(RPN []*structs.Token) (int64, error) {
	stack := structs.NewInt64Stack(0)
	for _, i := range RPN {
		switch i.Type {
		case structs.IsNum:
			stack.Push(i.Val.(int64))
		case structs.IsBinaryOperator:
			operand1, ok := stack.Pop()
			if !ok {
				return 0, ErrOperandMissed
			}
			operand2, ok := stack.Pop()
			if !ok {
				return 0, ErrOperandMissed
			}
			result, err := InvokeOperator(i, operand1, operand2)
			if err != nil {
				return 0, err
			}
			stack.Push(result)
		case structs.IsUnaryOperator:
			operand, ok := stack.Pop()
			if !ok {
				return 0, ErrOperandMissed
			}
			result, err := InvokeOperator(i, operand)
			if err != nil {
				return 0, err
			}
			stack.Push(result)
		default:
			return 0, structs.ErrBadToken
		}
	}
	ans, ok := stack.Pop()
	if !ok {
		return 0, ErrOperandMissed
	}
	if !stack.Empty() {
		return 0, ErrOperatorMissed
	}

	return ans, nil
}

func InvokeOperator(operator *structs.Token, operands ...int64) (int64, error) {
	if operator.Type == structs.IsUnaryOperator {
		if operator.Val.(rune) == '-' {
			return -1 * operands[0], nil
		}
		return operands[0], nil
	}

	switch c := operator.Val.(rune); c {
	case '+':
		return Add64(operands[0], operands[1])
	case '-':
		return Sub64(operands[1], operands[0])
	case '*':
		return Mult64(operands[0], operands[1])
	case '/':
		if operands[0] == 0 {
			return 0, ErrDivideByZero
		}
		return int64(math.Floor(float64(operands[1]) / float64(operands[0]))), nil
	default:
		return 0, ErrUnknownBinaryOperator
	}
}

func Add64(operand1, operand2 int64) (int64, error) {
	if operand2 >= 0 && operand1 > math.MaxInt64-operand2 || operand2 < 0 && operand1 < math.MinInt64-operand2 {
		return 0, ErrInt64Overflow
	}
	return operand1 + operand2, nil
}

func Sub64(operand1, operand2 int64) (int64, error) {
	return Add64(operand1, -1*operand2)
}

func Mult64(operand1, operand2 int64) (int64, error) {
	if operand1 == 0 || operand2 == 0 {
		return 0, nil
	}
	sgn1 := Sgn64(operand1)
	sgn2 := Sgn64(operand2)
	if sgn1*sgn2 == 1 {
		if operand1 == math.MinInt64 || operand2 == math.MinInt64 {
			return 0, ErrInt64Overflow
		}
		operand1 *= sgn1
		operand2 *= sgn2
		if operand1 > math.MaxInt64/operand2 {
			return 0, ErrInt64Overflow
		}
	} else {
		if Min64(operand1, operand2) < math.MinInt64/Max64(operand1, operand2) {
			return 0, ErrInt64Overflow
		}
	}
	return operand1 * operand2, nil
}

func Min64(operand1, operand2 int64) int64 {
	if operand1 < operand2 {
		return operand1
	}
	return operand2
}

func Max64(operand1, operand2 int64) int64 {
	if operand1 > operand2 {
		return operand1
	}
	return operand2
}

func Sgn64(operand int64) int64 {
	if operand >= 0 {
		return 1
	}
	return -1
}
