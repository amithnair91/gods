package set_test

import (
	"testing"

	s "github.com/amithnair91/gods/datastructures/set"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToCreateNewSet(t *testing.T) {
	set := s.NewSet()

	assert.IsType(t, s.Set{}, set)
}

func TestSizeReturnsZeroWhenSetIsEmpty(t *testing.T) {
	set := s.NewSet()

	assert.Equal(t, 0, set.Size())
}

func TestAddShouldNotPanic(t *testing.T) {
	set := s.NewSet()

	assert.NotPanics(t, func() { set.Add("someData") })
}

func TestSizeReturnsSizeOfSet(t *testing.T) {
	set := s.NewSet()

	set.Add("someData")

	assert.Equal(t, 1, set.Size())
}

func TestShouldReturnFalseWhenItemAlreadyExists(t *testing.T) {
	set := s.NewSet()

	set.Add("someData")
	err := set.Add("someData")

	assert.Error(t, err)
	assert.Equal(t, "cannot add item that already exists in the set", err.Error())
}
