package sorting

func BubbleSort(NumberList []int, SortOrder int) {
	n := len(NumberList)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1 - i; j++ {
			if SortOrder == 0 {
				if NumberList[j] > NumberList[j + 1] {
					NumberList[j + 1], NumberList[j] = NumberList[j], NumberList[j + 1]
				}
			} else if SortOrder == 1 {
				if NumberList[j] < NumberList[j + 1] {
					NumberList[j + 1], NumberList[j] = NumberList[j], NumberList[j + 1]
				}
			} else {
				panic("SortOrder can only have values - 0 and 1")
			} 
		}
	}
}

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