package stack_test

import (
	"testing"

	s "github.com/amithnair91/gods/datastructures/stack"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToCreateNewStack(t *testing.T) {
	stack := s.NewStack()

	assert.IsType(t, s.Stack{}, stack)
}
