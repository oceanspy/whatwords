package wordparser

import (
	"reflect"
	"testing"
)

type TestWordList struct {
	WordList []string
	Word     string
	Expected string
}

func TestSortByCount(t *testing.T) {
	wordList := []WordInfo{
		{Word: "enfin", Count: 1},
		{Word: "test", Count: 2},
		{Word: "voila", Count: 1},
		{Word: "simple", Count: 3},
	}

	expectedWordList := []WordInfo{
		{Word: "simple", Count: 3},
		{Word: "test", Count: 2},
		{Word: "enfin", Count: 1},
		{Word: "voila", Count: 1},
	}

	SortByCount(&wordList)

	for i := range wordList {
		if wordList[i] != expectedWordList[i] {
			t.Errorf("Expected sl doesn't match -- got: %v, expected: %v", wordList, expectedWordList)
		}
	}
}

func TestCalculateOccurenceOfEachWordInsideSliceBySorting(t *testing.T) {
	sl := []string{"simple", "test", "test", "enfin", "voila", "simple", "simple"}

	expectedWordList := []WordInfo{
		{Word: "simple", Count: 3},
		{Word: "test", Count: 2},
		{Word: "enfin", Count: 1},
		{Word: "voila", Count: 1},
	}

	wordList := CalculateOccurenceOfEachWordInsideSliceBySorting(&sl)
	SortByCount(&wordList)

	if !reflect.DeepEqual(wordList, expectedWordList) {
		t.Errorf("Expected wordList doesn't match -- got: %v, expected: %v", wordList, expectedWordList)
	}
}

func TestCalculateOccurenceOfEachWordInsideSliceByRemoving(t *testing.T) {
	sl := []string{"simple", "test", "test", "enfin", "voila", "simple", "simple"}

	expectedWordList := []WordInfo{
		{Word: "simple", Count: 3},
		{Word: "test", Count: 2},
		{Word: "enfin", Count: 1},
		{Word: "voila", Count: 1},
	}

	wordList := CalculateOccurenceOfEachWordInsideSliceByRemoving(&sl)
	SortByCount(&wordList)

	if !reflect.DeepEqual(wordList, expectedWordList) {
		t.Errorf("Expected wordList doesn't match -- got: %v, expected: %v", wordList, expectedWordList)
	}
}

func TestCountAndRemoveWord(t *testing.T) {
	sl := []string{"simple", "test", "test", "enfin", "voila", "simple", "simple"}

	word := "test"
	expectedCount := 2
	expectedSl := []string{"simple", "enfin", "voila", "simple", "simple"}
	count := CountAndRemoveWord(&sl, word)

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
		}
	}

	if count != expectedCount {
		t.Errorf("Expected count doesn't match -- got: %d, expected: %d", count, expectedCount)
	}

	word = "simple"
	expectedCount = 3
	expectedSl = []string{"enfin", "voila"}
	count = CountAndRemoveWord(&sl, word)

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
		}
	}

	if count != expectedCount {
		t.Errorf("Expected count doesn't match -- got: %d, expected: %d", count, expectedCount)
	}
}

func TestRemoveWord(t *testing.T) {
	sl := []string{"simple", "test", "test", "enfin", "voila", "simple", "simple"}
	word := "test"
	expectedSl := []string{"simple", "enfin", "voila", "simple", "simple"}

	RemoveWord(&sl, word)

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
		}
	}
}

func TestCountWord(t *testing.T) {
	sl := []string{"simple", "test", "test", "enfin", "voila", "simple", "simple"}
	word := "test"

	expectedCount := 2
	count := CountWord(&sl, word)

	if count != expectedCount {
		t.Errorf("Expected count doesn't match -- got: %d, expected: %d", count, expectedCount)
	}
}

func TestRemoveSliceElement(t *testing.T) {
	sl := []string{"simple", "test1", "test2", "enfin", "voila"}
	i := 2
	expectedSl := []string{"simple", "test2", "enfin", "voila"}

	RemoveSliceElement(&sl, i)

	if reflect.DeepEqual(sl, expectedSl) {
		t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
	}
}

func TestRemoveWordShorterThanExcept(t *testing.T) {
	sl := []string{"simple", "test", "a", "àà", "enfin", "voila", "simple", "simple", "e", "ee", "eee", "c", "cc", "ccc"}
	length := 2
	expectedSl := []string{"simple", "test", "enfin", "voila", "simple", "simple", "eee", "c"}
	RemoveWordShorterThanExcept(&sl, length)

	if reflect.DeepEqual(sl, expectedSl) {
		t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
	}
}

func TestRemoveSpecialCharacters(t *testing.T) {
	s := "éàçèù$-<>-$(`)[]{}||\\^@#~ôûêâôî"
	expectedS := "eaceuoueaoi"
	RemoveSpecialCharacters(&s)

	if s != expectedS {
		t.Errorf("Expected string doesn't match -- got: %s, expected: %s", s, expectedS)
	}
}
