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
