package sorting

// Performs bubble sort algorithm on the given number slice to sort the numbers as per the given sort order.
// Sort Order value of 0 implies ascending order.
// Sort Order value of 1 implies descending order.
func BubbleSort(NumberList []int, SortOrder int) {
	n := len(NumberList)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1 - i; j++ {
			if SortOrder == 0 && NumberList[j] > NumberList[j + 1]{
				NumberList[j + 1], NumberList[j] = NumberList[j], NumberList[j + 1]
			} else if SortOrder == 1 && NumberList[j] < NumberList[j + 1] {
				NumberList[j + 1], NumberList[j] = NumberList[j], NumberList[j + 1]
			} else {
				if SortOrder != 0 && SortOrder != 1 {
					panic("SortOrder can only have values - 0 and 1")
				}
			} 
		}
	}
}

// Performs insertion sort algorithm on the given number slice to sort the numbers as per the given sort order.
// Sort Order value of 0 implies ascending order.
// Sort Order value of 1 implies descending order.
func InsertionSort(NumberList []int, SortOrder int) {
	n := len(NumberList)
	for i := 1; i < n; i++ {
		refValue := NumberList[i]
		j := i - 1
		if SortOrder == 0 {
			for ; j >= 0 && NumberList[j] > refValue; j-- {
				NumberList[j + 1] = NumberList[j]
			}
		} else if SortOrder == 1 {
			for ; j >= 0 && NumberList[j] < refValue; j-- {
				NumberList[j + 1] = NumberList[j]
			}
		} else {
			panic("SortOrder can only have values - 0 and 1")
		}
		
		NumberList[j + 1] = refValue
	}
}

// Performs selection sort algorithm on the given number slice to sort the numbers as per the given sort order.
// Sort Order value of 0 implies ascending order.
// Sort Order value of 1 implies descending order.
func SelectionSort(NumberList []int, SortOrder int) {
	n := len(NumberList)
	for i := 0; i < n - 1; i++ {
		refValue := NumberList[i]
		index := i
		for j := i + 1; j < n; j++ {
			if SortOrder == 0 && NumberList[j] < refValue{
				refValue = NumberList[j]
				index = j
			} else if SortOrder == 1 && NumberList[j] > refValue {
				refValue = NumberList[j]
				index = j
			}
		}

		NumberList[i], NumberList[index] = NumberList[index], NumberList[i]
	}
}