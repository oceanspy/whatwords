package wordparser

import (
	"golang.org/x/text/unicode/norm"
	"regexp"
	"slices"
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

func RemoveWordShorterThanExcept(sl *[]string, length int, excludedShortWords []string) {
	var newSlice []string

	for _, str := range *sl {
		if len(str) <= length && !slices.Contains(excludedShortWords, str) {
			continue
		}
		newSlice = append(newSlice, str)
	}

	*sl = newSlice
}

func ReplaceSimilarWords(sl *[]string, similarWords map[string]string) {
	for i := range *sl {
		if _, ok := similarWords[(*sl)[i]]; ok {
			ReplaceSimilarWord(&(*sl)[i], similarWords)
		}
	}
}

func ReplaceSimilarWord(s *string, similarWords map[string]string) {
	if val, ok := similarWords[*s]; ok {
		*s = val
	}
}

func ReplaceMultipleWords(sl *[]string, multipleWords map[string]string) {
	var sliceItemToRemove []int
	for wordsOfMultipleWords, wordValue := range multipleWords {

		// Break the multiple words into a slice of words
		words := strings.Split(wordsOfMultipleWords, " ")

		for i := range *sl {
			// check if current sl word is the first word of the multiple words
			if (*sl)[i] != words[0] {
				continue
			}

			// Check if all the words are the same and in the right order
			success := true
			for j := range words {
				if words[j] != (*sl)[i+j] {
					success = false
					break
				}
			}

			if !success {
				continue
			}

			// Add the value to the slice
			*sl = append(*sl, wordValue)

			// Mark the words to remove from the slice
			for j := range words {
				sliceItemToRemove = append(sliceItemToRemove, i+j)
			}
		}
	}

	// Remove the words from the slice
	for i := len(sliceItemToRemove) - 1; i >= 0; i-- {
		RemoveSliceElement(sl, sliceItemToRemove[i])
	}
}

func RemoveExcludedWords(sl *[]string, excludedWords []string) {
	for _, word := range excludedWords {
		// RemoveWord(sl, word)
		CountAndRemoveWord(sl, word)
	}
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

func RemoveLineBreaks(sl *[]string) {
	for i := range *sl {
		if strings.Contains((*sl)[i], "\n") {
			(*sl)[i] = strings.ReplaceAll((*sl)[i], "\n", "")
		}
	}
}

func MakeLowerCaseFromList(sl *[]string) {
	for i := range *sl {
		MakeLowerCase(&(*sl)[i])
	}
}

func MakeLowerCase(s *string) {
	*s = strings.ToLower(*s)
}
