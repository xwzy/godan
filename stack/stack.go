package stack

type Stack struct {
	data []interface{}
}

func DefaultStack() *Stack {
	return NewStack()
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}

func (s *Stack) Push(elem interface{}) {
	s.data = append(s.data, elem)
}

func (s *Stack) Pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Top() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}
