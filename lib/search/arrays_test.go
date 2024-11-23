package search

import (
	"testing"
	"math/rand/v2"
)

func Test_LinearSearch(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	RandIndex := rand.IntN(N)
	SearchIndex := LinearSearch(NumberList, NumberList[RandIndex])
	if SearchIndex == -1 {
		t.Fatalf("Linear Search :: Search element - %d was not found in the number list.", NumberList[RandIndex])
	}
	
	if NumberList[RandIndex] == NumberList[SearchIndex] {
		t.Logf("Linear Search :: Search element - %d was successfully found in the number list.", NumberList[RandIndex])
	}
}

func Test_BinarySearch(t *testing.T) {
	NumberList := make([]int, 0)
	N := rand.IntN(100)
	t.Logf("NumberList to be generated will have %d numbers.", N)
	for i := 1; i <= N; i++ {
		randomNumber := rand.IntN(100)
		NumberList = append(NumberList, randomNumber)
	}

	RandIndex := rand.IntN(N)
	SearchIndex := BinarySearch(NumberList, NumberList[RandIndex])
	if SearchIndex == -1 {
		t.Fatalf("Binary Search :: Search Element - %d was not found in the number list.", NumberList[RandIndex])
	}
	
	if NumberList[RandIndex] == NumberList[SearchIndex] {
		t.Logf("Binary Search :: Search Element - %d was found successfully in the number list.", NumberList[RandIndex])
	}
}