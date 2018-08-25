package datastructures_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/amithnair/gods/datastructures"
	"reflect"
	"testing"
)

func TestShouldBeAbleToAddItem(t *testing.T) {
	linkedList := datastructures.NewSinglyLinkedList()
	added := linkedList.Add(1)

	assert.True(t, added)
}

func TestShouldAssignHeadElementOnFirstAdd(t *testing.T) {
	linkedList := datastructures.NewSinglyLinkedList()
	expectedHead := 1
	linkedList.Add(expectedHead)
	head := linkedList.Head()
	assert.Equal(t, expectedHead, head.Value())
}

func TestShouldReturnFirstAddedItemAsHead(t *testing.T) {

	inputItems := []string{"mamooty", "mohanlal"}

	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])

	head := linkedList.Head()

	assert.Equal(t, inputItems[0], head.Value())
}

func TestShouldBeAbleToReturnNextItem(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])

	head := linkedList.Head()
	nextItem := head.Next()

	assert.Equal(t, inputItems[1], nextItem.Value())
}

func TestShouldReturnTailAsEmptyOnAddingSingleElement(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	tail := linkedList.Tail()
	expectedResult := datastructures.Element{}
	assert.True(t, reflect.DeepEqual(expectedResult, tail))
}

func TestShouldReturnFalseOnRemoveElementIfItemDoesNotExist(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal", "srinivasan"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])
	linkedList.Add(inputItems[2])

	removed := linkedList.Remove("prithviraj")

	assert.False(t, removed)
}

func TestShouldReturnTrueIfElementExists(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal", "srinivasan"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])
	linkedList.Add(inputItems[2])

	removed := linkedList.Remove(inputItems[1])

	assert.True(t, removed)
}
