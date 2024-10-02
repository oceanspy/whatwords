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

func TestCalculateOccurenceOf(t *testing.T) {
	slToCount := []string{"simple", "test"}
	sl := []string{"simple", "test", "test", "enfin", "voila", "simple", "simple"}

	expectedWordList := []WordInfo{
		{Word: "simple", Count: 3},
		{Word: "test", Count: 2},
	}

	wordList := CalculateOccurenceOf(&slToCount, &sl)
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

func TestSplitWordsByDelimiter(t *testing.T) {
	sl := []string{"ra,simple,test;enfin!voila simple simple"}
	expectedSl := []string{"ra", "simple", "test", "enfin", "voila", "simple", "simple"}

	SplitWordsByDelimiters(&sl)

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl[i], expectedSl[i])
		}
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
	excludedShortWords := []string{"c"}

	sl := []string{"simple", "test", "a", "àà", "enfin", "voila", "simple", "simple", "e", "ee", "eee", "c", "cc", "ccc"}
	length := 2
	expectedSl := []string{"simple", "test", "enfin", "voila", "simple", "simple", "eee", "c"}
	RemoveWordShorterThanExcept(&sl, length, excludedShortWords)

	if reflect.DeepEqual(sl, expectedSl) {
		t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
	}
}

func TestReplaceSimilarWords(t *testing.T) {
	similarWords := map[string]string{
		"voila": "simple",
		"enfin": "test",
	}

	sl := []string{"simple", "test", "enfin", "voila", "simple", "simple"}
	ReplaceSimilarWords(&sl, similarWords)

	expectedSl := []string{"simple", "test", "test", "simple", "simple", "simple"}

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
		}
	}
}

func TestReplaceMultipleWords(t *testing.T) {
	multipleWords := map[string]string{
		"simple test": "simple-test",
	}

	sl := []string{"simple", "test", "enfin", "voila", "simple", "simple"}
	ReplaceMultipleWords(&sl, multipleWords)

	expectedSl := []string{"enfin", "voila", "simple", "simple", "simple-test"}

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
		}
	}
}

func TestRemoveExcludedWords(t *testing.T) {
	excludedWords := []string{"simple", "test"}

	sl := []string{"simple", "test", "enfin", "voila", "simple", "simple"}
	RemoveExcludedWords(&sl, excludedWords)

	expectedSl := []string{"enfin", "voila"}

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
		}
	}
}

func TestRemoveSpecialCharactersFromString(t *testing.T) {
	s := "éàçèù$-<>-$(`)[]{}||\\^@#~ôûêâôî"
	expectedS := "eaceuoueaoi"
	RemoveSpecialCharactersFromString(&s)

	if s != expectedS {
		t.Errorf("Expected string doesn't match -- got: %s, expected: %s", s, expectedS)
	}
}

func TestRemoveLineBreaks(t *testing.T) {
	sl := []string{"simple", "test\n", "enfin\n", "voila", "simple\n", "simple\n"}
	expectedSl := []string{"simple", "test", "enfin", "voila", "simple", "simple"}
	RemoveLineBreaks(&sl)

	for i := range sl {
		if sl[i] != expectedSl[i] {
			t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
		}
	}
}
