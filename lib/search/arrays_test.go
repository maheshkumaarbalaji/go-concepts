package search

import (
	"testing"
	"math/rand/v2"
	"strings"
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

func Test_Search(t *testing.T) {
	testCases := []struct {
		Name string
		SearchType string
	} {
		{ "Linear Search", "linear" },
		{ "Binary Search", "binary" },
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			NumberList := GenerateNumberList(t)
			N := len(NumberList)
			RandIndex := rand.IntN(N)
			SearchIndex := -1
			if strings.EqualFold(testCase.SearchType, "linear") {
				SearchIndex = LinearSearch(NumberList, NumberList[RandIndex])
			} else if strings.EqualFold(testCase.SearchType, "binary") {
				SearchIndex = BinarySearch(NumberList, NumberList[RandIndex])
			}

			if SearchIndex == -1 {
				t.Fatalf("%s :: Search element - %d was not found in the number list.", testCase.Name, NumberList[RandIndex])
			}

			if NumberList[RandIndex] == NumberList[SearchIndex] {
				t.Logf("%s:: Search element - %d was successfully found in the number list.", testCase.Name, NumberList[RandIndex])
			}
		})
	}
}