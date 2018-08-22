package datastructures_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gitlab.com/amithnair/gods/datastructures"
)

func TestShouldBeAbleToAddItem(t *testing.T) {
	linkedList := datastructures.NewLinkedList()
	added := linkedList.Add(1)

	assert.True(t,added)
}

func TestShouldAssignHeadElementOnFirstAdd(t *testing.T)  {
	linkedList := datastructures.NewLinkedList()
	expectedHead := 1
	linkedList.Add(expectedHead)
	head := linkedList.GetHead()
	assert.Equal(t,expectedHead,head)
}

func TestShouldReturnFirstAddedItemAsHead(t *testing.T) {

	inputItems := []string {"mamooty","mohanlal"}

	linkedList := datastructures.NewLinkedList()
	linkedList.Add(inputItems[0])
	linkedList.Add(inputItems[1])

	head := linkedList.GetHead()

	assert.Equal(t,inputItems[0],head)
}

