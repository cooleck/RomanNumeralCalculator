package structs

type Int64Stack []int64

func NewInt64Stack(n int) Int64Stack {
	return make(Int64Stack, 0, n)
}

func (s *Int64Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Int64Stack) Top() (val int64, ok bool) {
	if s.Empty() {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Int64Stack) Push(val int64) {
	*s = append(*s, val)
}

func (s *Int64Stack) Pop() (val int64, ok bool) {
	val, ok = s.Top()
	if ok {
		*s = (*s)[:len(*s)-1]
	}
	return val, ok
}
