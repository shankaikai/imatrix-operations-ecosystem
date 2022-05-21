// This file contains useful generic functions that could be used
// in any other packages.

package common

import "fmt"

const (
	CUSTOM_NEG_INF = -999999999999
)

// Checks if an element (key) exists within a sorted list
// Returns if the element is found and index of the key
// in the array if found. If the key is not found, the index
// will be custom -inf for
func BinarySearch(list []int, left int, right int, key int) (bool, int) {
	fmt.Println("Binary Search", left, right, key, list)
	if right < left {
		return false, CUSTOM_NEG_INF
	}

	mid := left + (right-left)/2

	if list[mid] == key {
		return true, mid
	}
	if key < list[mid] {
		return BinarySearch(list, left, mid, key)
	} else {
		return BinarySearch(list, mid+1, right, key)
	}
}
