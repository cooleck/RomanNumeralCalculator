package algorithm

import (
	"errors"
	"math"
	"strings"
	"vk/structs"
)

var (
	nums = map[rune]int64{
		'Z': 0,
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	decNums = map[rune]struct{}{
		'I': {},
		'X': {},
		'C': {},
	}

	symbols = map[rune]struct{}{
		'+': {},
		'-': {},
		'*': {},
		'/': {},
		'(': {},
		')': {},
	}

	numsComb = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	numsCombKeys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	ErrIncorrectNumber = errors.New("incorrect number")
	ErrUnknownSymbol   = errors.New("unknown symbol")
	ErrRomanOverflow   = errors.New("the result is too much for Roman numeral system")
)

func ParseInput(input string) ([]*structs.Token, error) {

	ans := make([]*structs.Token, 0)
	isNum := false
	isPrevNum := false
	num := strings.Builder{}

	for _, i := range input {
		if i == ' ' {
			continue
		}
		if _, ok := symbols[i]; ok {
			if isNum {
				n, err := ConvertToArabic(num.String())
				if err != nil {
					return nil, err
				}

				ans = append(ans, structs.NewToken(0, n, structs.IsNum))

				num.Reset()
				isNum = false
				isPrevNum = true
			}

			token := &structs.Token{Val: i, Precedence: 0}

			switch i {
			case '+':
				token.Type = structs.IsBinaryOperator
				token.Precedence = 1
				isPrevNum = false
			case '-':
				if isPrevNum {
					token.Type = structs.IsBinaryOperator
					token.Precedence = 1
					isPrevNum = false
				} else {
					token.Type = structs.IsUnaryOperator
					token.Precedence = 3
				}
			case '*', '/':
				token.Type = structs.IsBinaryOperator
				token.Precedence = 2
				isPrevNum = false
			case '(':
				token.Type = structs.IsOpenBracket
				isPrevNum = false
			case ')':
				token.Type = structs.IsCloseBracket
			}

			ans = append(ans, token)
		} else if _, ok = nums[i]; ok {
			if !isNum {
				isNum = true
			}
			num.WriteRune(i)
		} else {
			return nil, ErrUnknownSymbol
		}
	}

	if isNum {
		n, err := ConvertToArabic(num.String())
		if err != nil {
			return nil, err
		}
		ans = append(ans, structs.NewToken(0, n, structs.IsNum))
	}

	return ans, nil
}

func isCorrect(n string) bool {
	m := make(map[rune]int)
	isPrevDec := false
	var upper int64

	for i, j := range n {
		if j == 'Z' && i != 0 {
			return false
		}
		if i == 0 {
			if _, ok := decNums[j]; ok {
				upper = 10*nums[j] + 1
				isPrevDec = true
			} else if j == 'M' {
				upper = nums[j] + 1
			} else {
				upper = nums[j]
			}
			m[j]++
		} else if nums[j] < upper {
			if isPrevDec {
				if nums[j] == nums[rune(n[i-1])]*10 || nums[j] == nums[rune(n[i-1])]*5 {
					upper = nums[rune(n[i-1])]
				} else {
					if _, ok := decNums[j]; ok {
						if rune(n[i-1]) == j {
							if v, _ := m[j]; v < 3 {
								m[j]++
								upper = nums[j] + 1
							} else {
								return false
							}
						} else {
							if v, _ := m[j]; v < 3 {
								m[j]++
								upper = 10*nums[j] + 1
								isPrevDec = true
							} else {
								return false
							}
						}
					} else {
						upper = nums[j]
					}

				}
			} else if _, ok := decNums[j]; ok {
				upper = 10*nums[j] + 1
				isPrevDec = true
			} else if j == 'M' {
				if v, _ := m[j]; v < 3 {
					m[j]++
					upper = nums[j] + 1
				} else {
					return false
				}
			} else {
				upper = nums[j]
			}
		} else {
			return false
		}
	}
	return true
}

func ConvertToArabic(n string) (int64, error) {
	if !isCorrect(n) {
		return 0, ErrIncorrectNumber
	}

	var ans int64 = 0

	for i := 0; i < len(n)-1; i++ {
		if nums[rune(n[i])] >= nums[rune(n[i+1])] {
			ans += nums[rune(n[i])]
		} else {
			ans -= nums[rune(n[i])]
		}
	}

	ans += nums[rune(n[len(n)-1])]

	return ans, nil
}

func ConvertToRoman(n int64) (string, error) {
	if n == 0 {
		return "Z", nil
	}

	sgn := ""
	if n < 0 {
		sgn = "-"
	}
	n = int64(math.Abs(float64(n)))
	if n > 3999 {
		return "", ErrRomanOverflow
	}

	num := int(n)

	ans := strings.Builder{}
	ans.WriteString(sgn)

	for _, i := range numsCombKeys {
		ans.WriteString(strings.Repeat(numsComb[i], num/i))
		num %= i
	}

	return ans.String(), nil
}
