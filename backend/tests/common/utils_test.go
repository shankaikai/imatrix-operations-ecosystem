package tests

import (
	"testing"

	"capstone.operations_ecosystem/backend/common"
	"github.com/stretchr/testify/assert"
)

// First in list
func TestBinarySearchFirst(t *testing.T) {
	list := []int{4, 5, 7, 8, 10, 78}
	found, index := common.BinarySearch(list, 0, 5, 4)
	assert.Equal(t, true, found, "The number should be in the list")
	assert.Equal(t, 0, index, "The number should be in the first in the list")
}

// Last in list
func TestBinarySearchLast(t *testing.T) {
	list := []int{4, 5, 7, 8, 10, 78}
	found, index := common.BinarySearch(list, 0, 5, 78)
	assert.Equal(t, true, found, "The number should be in the list")
	assert.Equal(t, 5, index, "The number should be in the 5th in the list")
}

// middle of list
func TestBinarySearchMiddle(t *testing.T) {
	list := []int{4, 5, 7, 8, 10, 78}
	found, index := common.BinarySearch(list, 0, 5, 8)
	assert.Equal(t, true, found, "The number should be in the list")
	assert.Equal(t, 3, index, "The number should be in the 3rd in the list")
}

// not in list
func TestBinarySearchNotInList(t *testing.T) {
	list := []int{4, 5, 7, 8, 10, 78}
	found, index := common.BinarySearch(list, 0, 5, 4324)
	assert.Equal(t, false, found, "The number should be in the list")
	assert.Equal(t, common.CUSTOM_NEG_INF, index, "The number should not be in the list")
}
