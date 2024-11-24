package search

import (
	"github.com/maheshkumaarbalaji/goconcepts/lib/sorting"
)

type SortType interface {
	int | float32 | float64
}

// Uses Linear Search algorithm to determine the index of the first occurrence of the search element in the given number slice.
func LinearSearch[T SortType](NumberList []T, SearchNum T) int {
	TargetElemPos := -1

	for i, elem := range NumberList {
		if elem == SearchNum {
			TargetElemPos = i
			break
		}
	}

	return TargetElemPos
}

// Uses Binary Search algorithm to determine the index of the first occurrence of the search element in the given number slice.
func BinarySearch[T SortType](NumberList []T, SearchNum T) int {
	TargetElemPos := -1
	l, n := 0, len(NumberList) - 1
	err := sorting.BubbleSort(NumberList, 0)
	if err != nil {
		return -1
	}
	
	for ; l <= n; {
		mid := (l + n)/2
		if NumberList[mid] == SearchNum {
			TargetElemPos = mid
			break
		} else if SearchNum > NumberList[mid] {
			l = mid + 1
		} else {
			n = mid - 1
		}
	}

	return TargetElemPos
}