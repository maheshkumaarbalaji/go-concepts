package search

import (
	"github.com/maheshkumaarbalaji/goconcepts/lib/sorting"
)

func LinearSearch(NumberList []int, SearchNum int) int {
	TargetElemPos := -1

	for i, elem := range NumberList {
		if elem == SearchNum {
			TargetElemPos = i
			break
		}
	}

	return TargetElemPos
}

func BinarySearch(NumberList []int, SearchNum int) int {
	TargetElemPos := -1
	l, n := 0, len(NumberList) - 1
	sorting.BubbleSort(NumberList, 0)
	
	for ; l < n; {
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