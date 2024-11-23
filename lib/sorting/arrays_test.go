package sorting

import (
	"math/rand/v2"
	"strings"
	"testing"
)

func GenerateNumberList(t testing.TB) []int {
	t.Helper()
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("Number list to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	return NumberList
}

func IsSorted(t testing.TB, NumberList []int, SortOrder int) (bool, int) {
	t.Helper()
	isSorted := true
	errIndex := -1
	N := len(NumberList)
	for index := 1; index < N; index++ {
		if (SortOrder == 0 && NumberList[index] < NumberList[index - 1]) || (SortOrder == 1 && NumberList[index] > NumberList[index - 1]) {
			isSorted = false
			errIndex = index
			break
		}
	}

	return isSorted, errIndex
}

func Test_Sorting(t *testing.T) {
	testCases := []struct {
		Name string
		SortingType string
		SortOrder int
	} {
		{ "Bubble Sort - Ascending Order", "bubble", 0 },
		{ "Bubble Sort - Descending Order", "bubble", 1 },
		{ "Insertion Sort - Ascending Order", "insertion", 0 },
		{ "Insertion Sort - Descending Order", "insertion", 1 },
		{ "Selection Sort - Ascending Order", "selection", 0 },
		{ "Selection Sort - Descending Order", "selection", 1 },
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			NumberList := GenerateNumberList(tt)
			tt.Log("Number list generated for testing: ", NumberList)
			var err error
			if strings.EqualFold(testCase.SortingType, "bubble") {
				err = BubbleSort(NumberList, testCase.SortOrder)
			} else if strings.EqualFold(testCase.SortingType, "insertion") {
				err = InsertionSort(NumberList, testCase.SortOrder)
			} else if strings.EqualFold(testCase.SortingType, "selection") {
				err = SelectionSort(NumberList, testCase.SortOrder)
			}

			if err != nil {
				tt.Fatalf("%s :: Ascending Order :: Sorting process failed with message: %s", testCase.Name, err.Error())
				return
			}

			tt.Log("Number list after sorting: ", NumberList)
			isSorted, errIndex := IsSorted(tt, NumberList, testCase.SortOrder)
			if !isSorted {
				tt.Fatalf("%s :: Sorting process failed - Number at index %d not sorted", testCase.Name, errIndex)
			} else {
				tt.Logf("%s :: Sorting process successfull", testCase.Name)
			}
		})
	}
}