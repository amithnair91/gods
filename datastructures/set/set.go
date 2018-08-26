package set

import (
	"errors"
)

type Set struct {
	items []interface{}
}

func NewSet() Set {
	return Set{}
}

func (s *Set) Size() int {
	return len(s.items)
}

//time complexity O(n) worst case
func (s *Set) Remove(item interface{}) bool {

	if index, exists := s.contains(item); exists {
		var firstSlice []interface{}
		var secondSlice []interface{}

		if index == 0 {
			firstSlice = s.items[1:]
		} else {
			firstSlice = s.items[0:index]
		}

		if index >= s.Size() {
			secondSlice = s.items[:s.Size()-1]
		} else {
			secondSlice = s.items[index+1 : s.Size()]
		}

		s.items = append(firstSlice, secondSlice...)
		return true
	}

	return false
}

//time complexity O(n) worst case
func (s *Set) Add(item interface{}) error {

	if _, exists := s.contains(item); exists {
		return errors.New("cannot add item that already exists in the set")
	}

	s.items = append(s.items, item)
	return nil
}

func (s *Set) contains(item interface{}) (int, bool) {
	for index, val := range s.items {
		if val == item {
			return index, true
		}
	}
	return 0, false
}

func (s *Set) Contains(item interface{}) bool {

	if _, exists := s.contains(item); exists {
		return true
	}

	return false
}
