package stack

import "errors"

type Stack struct {
	items []interface{}
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Peek() interface{} {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() (interface{}, error) {

	if len(s.items) < 1 {
		return nil, errors.New("cannot pop as stack is empty")
	}

	poppedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return poppedItem, nil
}
