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

func CalculateOccurenceOf(slToCount *[]string, sl *[]string) []WordInfo {
	var wordList []WordInfo

	for _, wordToCount := range *slToCount {
		count := CountWord(sl, wordToCount)
		wordList = append(wordList, WordInfo{Word: wordToCount, Count: count})
	}

	return wordList
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

func SplitWordsByDelimiters(sl *[]string) {
	var delimiters = []string{",", ";", ":", ".", "!", "?", " "}

	for _, delimiter := range delimiters {
		SplitWordsByDelimiter(sl, delimiter)
	}
}

func SplitWordsByDelimiter(sl *[]string, delimiter string) {
	var sliceItemToRemove []int
	var toAppend []string

	for wordIndex, wordValue := range *sl {

		// If wordValue doesn't contain the delimiter, skip
		if !strings.Contains(wordValue, delimiter) {
			continue
		}

		// Break the word into a slice of words
		words := strings.Split(wordValue, delimiter)

		// Add the words to the slice
		toAppend = append(toAppend, words...)

		// Mark the word to remove from the slice
		sliceItemToRemove = append(sliceItemToRemove, wordIndex)
	}

	// Remove the words from the slice
	for i := len(sliceItemToRemove) - 1; i >= 0; i-- {
		RemoveSliceElement(sl, sliceItemToRemove[i])
	}

	// Append the new words to the slice
	*sl = append(*sl, toAppend...)
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

func RemoveEmptyWords(sl *[]string) {
	RemoveWord(sl, "")
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
	for i := len(*sl) - 1; i >= 0; i-- {
		if len((*sl)[i]) <= length && !slices.Contains(excludedShortWords, (*sl)[i]) {
			RemoveSliceElement(sl, i)
		}
	}
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

func RemoveSpecialCharacters(sl *[]string) {
	for i := range *sl {
		RemoveSpecialCharactersFromString(&(*sl)[i])
	}
}

func RemoveSpecialCharactersFromString(s *string) {
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

func MakeLowerCase(sl *[]string) {
	for i := range *sl {
		(*sl)[i] = strings.ToLower((*sl)[i])
	}
}
