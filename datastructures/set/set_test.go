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

	assert.NotPanics(t, func() { set.Add("someItem") })
}

func TestSizeReturnsSizeOfSet(t *testing.T) {
	set := s.NewSet()

	set.Add("someItem")

	assert.Equal(t, 1, set.Size())
}

func TestShouldReturnFalseWhenItemAlreadyExists(t *testing.T) {
	set := s.NewSet()

	set.Add("someItem")
	err := set.Add("someItem")

	assert.Error(t, err)
	assert.Equal(t, "cannot add item that already exists in the set", err.Error())
}

func TestRemoveReturnsFalseWhenItemDoesNotExists(t *testing.T) {
	set := s.NewSet()

	removed := set.Remove("someItem")

	assert.False(t, removed)
}

func TestRemoveReturnsIfItemExists(t *testing.T) {
	set := s.NewSet()

	set.Add("someItem")
	removed := set.Remove("someItem")

	assert.True(t, removed)
}

func TestShouldReduceSizeOnRemovalOfItem(t *testing.T) {
	set := s.NewSet()

	set.Add("firstItem")
	set.Add("secondItem")
	set.Add("thirdItem")

	assert.Equal(t, 3, set.Size())

	set.Remove("secondItem")
	assert.Equal(t, 2, set.Size())

	set.Remove("thirdItem")
	assert.Equal(t, 1, set.Size())

	set.Remove("firstItem")
	assert.Equal(t, 0, set.Size())
}
