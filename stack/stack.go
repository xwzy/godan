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

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Push(elem interface{}) {
	s.data = append(s.data, elem)
}

func (s *Stack) Pop() {
	if !s.Empty() {
		s.data = s.data[:len(s.data)-1]
	}
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[len(s.data)-1]
}
