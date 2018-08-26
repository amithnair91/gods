package datastructures_test

import (
	"reflect"
	"testing"

	"github.com/amithnair91/gods/datastructures"

	"github.com/stretchr/testify/assert"
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
	inputItems := []string{"mamooty"}
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

func TestShouldBeAbleToAddMultipleItems(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal", "srinivasan", "nivin"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])
	linkedList.Add(inputItems[2])
	linkedList.Add(inputItems[3])

	mamooty := linkedList.Head()
	mohanlal := mamooty.Next()
	srinivasan := mohanlal.Next()
	nivin := srinivasan.Next()

	assert.Equal(t, inputItems[0], mamooty.Value())
	assert.Equal(t, inputItems[1], mohanlal.Value())
	assert.Equal(t, inputItems[2], srinivasan.Value())
	assert.Equal(t, inputItems[3], nivin.Value())

}

func TestShouldBeAbleToListItems(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal", "srinivasan", "nivin"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])
	linkedList.Add(inputItems[2])
	linkedList.Add(inputItems[3])
	linkedList.Add(1)

	items := linkedList.Array()

	assert.Equal(t, inputItems[0], items[0])
	assert.Equal(t, inputItems[1], items[1])
	assert.Equal(t, inputItems[2], items[2])
	assert.Equal(t, inputItems[3], items[3])
	assert.Equal(t, 1, items[4])
}

func TestShouldBeAbleToAddItemsOfDifferentDataTypes(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal", "srinivasan", "nivin"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])
	linkedList.Add(inputItems[2])
	linkedList.Add(inputItems[3])
	linkedList.Add(1)

	items := linkedList.Array()

	assert.Equal(t, inputItems[0], items[0])
	assert.Equal(t, inputItems[1], items[1])
	assert.Equal(t, inputItems[2], items[2])
	assert.Equal(t, inputItems[3], items[3])
	assert.Equal(t, 1, items[4])
}

func TestShouldBeAbleToMaintainInsertionOrderOnRemoval(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal", "srinivasan", "nivin"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])
	linkedList.Add(inputItems[2])
	linkedList.Add(inputItems[3])

	linkedList.Remove(inputItems[1])

	mamooty := linkedList.Head()
	srinivasan := mamooty.Next()
	nivin := srinivasan.Next()

	assert.Equal(t, inputItems[0], mamooty.Value())
	assert.Equal(t, inputItems[2], srinivasan.Value())
	assert.Equal(t, inputItems[3], nivin.Value())

}

func TestShouldBeAbleToMaintainInsertionOrderOnRemovalOfMultipleElements(t *testing.T) {
	inputItems := []string{"mamooty", "mohanlal", "srinivasan"}
	linkedList := datastructures.NewSinglyLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])
	linkedList.Add(inputItems[2])
	linkedList.Add(1)
	linkedList.Add(2)

	linkedList.Remove(inputItems[1])
	linkedList.Remove(1)

	mamooty := linkedList.Head()
	srinivasan := mamooty.Next()
	number := srinivasan.Next()

	assert.Equal(t, inputItems[0], mamooty.Value())
	assert.Equal(t, inputItems[2], srinivasan.Value())
	assert.Equal(t, 2, number.Value())
}
