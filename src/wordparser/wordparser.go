package wordparser

import (
	"golang.org/x/text/unicode/norm"
	"regexp"
	"slices"
	"sort"
	"strings"
	"unicode"
)

var ExcludedShortWords = []string{"c"}
var WordsToReplace = map[string]string{
	"c++":         "cpp",
	"c#":          "csharp",
	"objective-c": "objectivec",
}
var WordsToMerge = map[string]string{
	"full stack": "fullstack",
	"front end":  "frontend",
	"back end":   "backend",
}

type WordInfo struct {
	Word  string
	Count int
}

func SortByCount(sl *[]WordInfo) {
	sort.Slice(*sl, func(i, j int) bool {
		return (*sl)[i].Count > (*sl)[j].Count
	})
}

func CalculateOccurenceOfEachWordInsideSliceBySorting(sl *[]string) []WordInfo {
	var wordList []WordInfo

	// Sort the slice
	sort.Strings(*sl)

	count := 0
	currentWord := (*sl)[0]
	for i := 0; i <= len(*sl); i++ {
		if i == len(*sl) {
			wordList = append(wordList, WordInfo{Word: currentWord, Count: count})
			break
		}

		if (*sl)[i] == currentWord {
			count++
			continue
		}

		wordList = append(wordList, WordInfo{Word: currentWord, Count: count})

		count = 1
		currentWord = (*sl)[i]
	}

	return wordList
}

func CalculateOccurenceOfEachWordInsideSliceByRemoving(sl *[]string) []WordInfo {
	var wordList []WordInfo

	for len(*sl) > 0 {
		word := (*sl)[0]
		count := CountAndRemoveWord(sl, word)
		wordList = append(wordList, WordInfo{Word: word, Count: count})
	}

	return wordList
}

func CalculateOccurenceOfEachWordInsideSlice(sl *[]string) []WordInfo {
	return CalculateOccurenceOfEachWordInsideSliceBySorting(sl)
}

func CountAndRemoveWord(sl *[]string, word string) int {
	count := CountWord(sl, word)
	RemoveWordAppearances(sl, word, count)

	return count
}

func CountWord(sl *[]string, word string) int {
	count := 0
	for i := range *sl {
		if (*sl)[i] == word {
			count++
		}
	}

	return count
}

func RemoveWordAppearances(sl *[]string, word string, appearances int) {
	for appearances > 0 {
		RemoveSliceElement(sl, slices.Index(*sl, word))
		appearances--
	}
}

func RemoveWord(sl *[]string, word string) {
	for slices.Contains(*sl, word) {
		RemoveSliceElement(sl, slices.Index(*sl, word))
	}
}

func RemoveSliceElement(sl *[]string, i int) {
	*sl = slices.Delete(*sl, i, i+1)
}

func RemoveWordShorterThanExcept(sl *[]string, length int) {
	var newSlice []string

	for _, str := range *sl {
		if len(str) < length && !slices.Contains(ExcludedShortWords, str) {
			continue
		}
		newSlice = append(newSlice, str)
	}

	*sl = newSlice
}

func RemoveSpecialCharactersFromList(sl *[]string) {
	for i := range *sl {
		RemoveSpecialCharacters(&(*sl)[i])
	}
}

func RemoveSpecialCharacters(s *string) {
	// Normalize the string (NFKD normal form decomposes characters)
	t := norm.NFKD.String(*s)

	// Remove non-ASCII characters using a custom filter
	filtered := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == ' ' {
			return r
		}
		return -1
	}, t)

	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)

	finalStr := re.ReplaceAllString(filtered, " ")

	*s = finalStr
}

func MakeLowerCaseFromList(sl *[]string) {
	for i := range *sl {
		MakeLowerCase(&(*sl)[i])
	}
}

func MakeLowerCase(s *string) {
	*s = strings.ToLower(*s)
}
