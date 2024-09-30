package wordparser

import (
	"golang.org/x/text/unicode/norm"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

type WordInfo struct {
	Word  string
	Count int
}

func SortByCount(sl *[]WordInfo) {
	sort.Slice(*sl, func(i, j int) bool {
		return (*sl)[i].Count > (*sl)[j].Count
	})
}

func CalculateOccurence(sl []string) []WordInfo {
	var wordList []WordInfo

	for len(sl) > 0 {
		word := sl[0]
		count := CountAndRemoveWord(&sl, word)
		wordList = append(wordList, WordInfo{Word: word, Count: count})
	}

	return wordList
}

func CountAndRemoveWord(sl *[]string, word string) int {
	count := 0
	for i, e := range *sl {
		if e == word {
			count++
			RemoveSliceElement(sl, i)
		}
	}

	return count
}

func RemoveSliceElement(sl *[]string, i int) {
	if i < 0 || i >= len(*sl) {
		return // Index out of bounds, do nothing
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
}

func RemoveUnwantedWords(sl *[]string) {
	var newSlice []string

	for _, str := range *sl {
		if len(str) < 2 {
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
