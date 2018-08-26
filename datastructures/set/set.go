package set

import "errors"

type Set struct {
	items []interface{}
}

func NewSet() Set {
	return Set{}
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) Add(item interface{}) error {

	if s.contains(item) {
		return errors.New("cannot add item that already exists in the set")
	}

	s.items = append(s.items, item)
	return nil
}

func (s *Set) contains(item interface{}) bool {
	for _, val := range s.items {
		if val == item {
			return true
		}
	}
	return false
}
