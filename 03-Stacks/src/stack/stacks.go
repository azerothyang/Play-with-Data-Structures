package stack

type Stack struct {
	Data  []interface{}
	Index int
}

// parentheses match cases
var parentheses = map[int32]int32{
	rune('{') : rune('}'),
	rune('[') : rune(']'),
	rune('(') : rune(')'),
}

func New(cap int) *Stack {
	return &Stack{
		Data:  make([]interface{}, cap),
		Index: 0,
	}
}

func (s *Stack) Push(val interface{}) {
	if len(s.Data) == s.Index {
		s.resize(s.Index * 2)
	}
	s.Data[s.Index] = val
	s.Index++
}

func (s *Stack) Pop() interface{} {
	// is stack is empty then return nil
	if s.IsEmpty() {
		return nil
	}
	if len(s.Data) == s.Index/4 {
		s.resize(s.Index / 2)
	}
	s.Index--
	return s.Data[s.Index]
}

func (s *Stack) Peek() interface{} {
	return s.Data[s.Index-1]
}

func (s *Stack) GetIndex() int {
	return s.Index
}

func (s *Stack) IsEmpty() bool {
	return s.Index == 0
}

func (s *Stack) resize(newCap int) {
	newS := New(newCap)
	// copy old stack to new stack
	for i := 0; i < s.Index; i++ {
		newS.Data[i] = s.Data[i]
	}
	s.Data = newS.Data
}

// check valid parentheses
func CheckValidParentheses(str string) bool {
	stacker := New(2)
	for _, ch := range str {
		if ch == rune('[') || ch == rune('(') || ch == ('{') {
			stacker.Push(ch)
		} else if ch == rune(']') || ch == rune(')') || ch == rune('}') {
			if stacker.IsEmpty() || parentheses[stacker.Pop().(int32)] != ch {
				return false
			}
		}
	}
	return stacker.IsEmpty()
}
