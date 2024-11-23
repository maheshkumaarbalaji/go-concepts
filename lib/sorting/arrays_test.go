package sorting

import (
	"testing"
	"math/rand/v2"
)

func Test_BubbleSort_Asc(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	err := BubbleSort(NumberList, 0)
	if err != nil {
		t.Fatalf("Bubble Sort - Ascending Order :: Sorting process failed with message: %s", err.Error())
		return
	}

	isSorted := true
	errIndex := -1
	for index := 1; index < N; index++ {
		if NumberList[index] < NumberList[index - 1] {
			isSorted = false
			errIndex = index
			break
		}
	}

	if !isSorted {
		t.Fatalf("Bubble Sort - Ascending Order :: Sorting process failed - Number at index %d not sorted", errIndex)
	} else {
		t.Log("Bubble Sort - Ascending Order :: Sorting process successfull")
	}
}

func Test_BubbleSort_Desc(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	err := BubbleSort(NumberList, 1)
	if err != nil {
		t.Fatalf("Bubble Sort - Descending Order :: Sorting process failed with message: %s", err.Error())
		return
	}

	isSorted := true
	errIndex := -1
	for index := 1; index < N; index++ {
		if NumberList[index] > NumberList[index - 1] {
			isSorted = false
			errIndex = index
			break
		}
	}

	if !isSorted {
		t.Fatalf("Bubble Sort - Descending Order :: Sorting process failed - Number at index %d not sorted", errIndex)
	} else {
		t.Log("Bubble Sort - Descending Order :: Sorting process successfull")
	}
}

func Test_InsertionSort_Asc(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	err := InsertionSort(NumberList, 0)
	if err != nil {
		t.Fatalf("Insertion Sort - Ascending Order :: Sorting process failed with message: %s", err.Error())
		return
	}

	isSorted := true
	errIndex := -1
	for index := 1; index < N; index++ {
		if NumberList[index] < NumberList[index - 1] {
			isSorted = false
			errIndex = index
			break
		}
	}

	if !isSorted {
		t.Fatalf("Insertion Sort - Ascending Order :: Sorting process failed - Number at index %d not sorted", errIndex)
	} else {
		t.Log("Insertion Sort - Ascending Order :: Sorting process successfull")
	}
}

func Test_InsertionSort_Desc(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	err := InsertionSort(NumberList, 1)
	if err != nil {
		t.Fatalf("Insertion Sort - Descending Order :: Sorting process failed with message: %s", err.Error())
		return
	}

	isSorted := true
	errIndex := -1
	for index := 1; index < N; index++ {
		if NumberList[index] > NumberList[index - 1] {
			isSorted = false
			errIndex = index
			break
		}
	}

	if !isSorted {
		t.Fatalf("Insertion Sort - Descending Order :: Sorting process failed - Number at index %d not sorted", errIndex)
	} else {
		t.Log("Insertion Sort - Descending Order :: Sorting process successfull")
	}
}

func Test_SelectionSort_Asc(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	err := SelectionSort(NumberList, 0)
	if err != nil {
		t.Fatalf("Selection Sort - Ascending Order :: Sorting process failed with message: %s", err.Error())
		return
	}

	isSorted := true
	errIndex := -1
	for index := 1; index < N; index++ {
		if NumberList[index] < NumberList[index - 1] {
			isSorted = false
			errIndex = index
			break
		}
	}

	if !isSorted {
		t.Fatalf("Selection Sort - Ascending Order :: Sorting process failed - Number at index %d not sorted", errIndex)
	} else {
		t.Log("Selection Sort - Ascending Order :: Sorting process successfull")
	}
}

func Test_SelectionSort_Desc(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	err := SelectionSort(NumberList, 1)
	if err != nil {
		t.Fatalf("Selection Sort - Descending Order :: Sorting process failed with message: %s", err.Error())
		return
	}

	isSorted := true
	errIndex := -1
	for index := 1; index < N; index++ {
		if NumberList[index] > NumberList[index - 1] {
			isSorted = false
			errIndex = index
			break
		}
	}

	if !isSorted {
		t.Fatalf("Selection Sort - Descending Order :: Sorting process failed - Number at index %d not sorted", errIndex)
	} else {
		t.Log("Selection Sort - Descending Order :: Sorting process successfull")
	}
}