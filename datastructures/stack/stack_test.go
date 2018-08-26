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

func TestPushItemShouldNotPanic(t *testing.T) {
	stack := s.NewStack()

	assert.NotPanics(t, func() { stack.Push("firstItem") })
}

func TestSizeReturnsZeroWhenStackIsEmpty(t *testing.T) {
	stack := s.NewStack()

	assert.Equal(t, 0, stack.Size())
}

func TestShouldBeAbleToGetStackSize(t *testing.T) {
	stack := s.NewStack()

	stack.Push("firstItem")
	stack.Push("secondItem")

	assert.Equal(t, 2, stack.Size())
}

func TestShouldBeAbleToPeekTopMostItem(t *testing.T) {
	expectedResult := "secondItem"
	stack := s.NewStack()

	stack.Push("firstItem")
	stack.Push(expectedResult)

	topMostItem := stack.Peek()

	assert.Equal(t, expectedResult, topMostItem)
}

func TestPopreturnsErrorWhenStackIsEmpty(t *testing.T) {
	stack := s.NewStack()

	_, error := stack.Pop()

	assert.Error(t, error)
	assert.Equal(t, "cannot pop as stack is empty", error.Error())
}

func TestPopReturnsNilErrorWhenPopIsSuccessful(t *testing.T) {
	stack := s.NewStack()

	stack.Push("firstItem")
	_, error := stack.Pop()

	assert.NoError(t, error)
}

func TestPopRemovesTopMostItem(t *testing.T) {
	expectedResult := "topMostItem"
	expectedSize := 2
	stack := s.NewStack()

	stack.Push("firstItem")
	stack.Push("secondItem")
	stack.Push(expectedResult)

	poppedItem, _ := stack.Pop()
	stackSize := stack.Size()

	assert.Equal(t, expectedResult, poppedItem)
	assert.Equal(t, expectedSize, stackSize)
}

func TestShouldBeAbleToPushItemsOfDifferentTypes(t *testing.T) {
	stringType := "firstItem"
	intType := 1
	floatType := 3.14
	stack := s.NewStack()

	stack.Push(stringType)
	stack.Push(intType)
	stack.Push(floatType)

	assert.Equal(t, 3, stack.Size())

	thirdItem, err := stack.Pop()
	assert.NoError(t, err)
	secondItem, err := stack.Pop()
	assert.NoError(t, err)
	firstItem, err := stack.Pop()
	assert.NoError(t, err)

	assert.Equal(t, floatType, thirdItem)
	assert.Equal(t, intType, secondItem)
	assert.Equal(t, stringType, firstItem)
	assert.Equal(t, 0, stack.Size())
}
