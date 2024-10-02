package wordlist

import (
	"whatwords/src/csvservice"
)

func GetExcludedWords() []string {
	excludedWords, err := csvservice.ToSlice("excludedwords.csv")
	if err != nil {
		return DefaultExcludedWords
	}

	if len(excludedWords) == 0 {
		return DefaultExcludedWords
	}

	return excludedWords
}

func GetShortWordsToKeep() []string {
	shortWordsToKeep, err := csvservice.ToSlice("shortwordstokeep.csv")
	if err != nil {
		return DefaultShortWordsToKeep
	}

	if len(shortWordsToKeep) == 0 {
		return DefaultShortWordsToKeep
	}

	return shortWordsToKeep
}

func GetMultipleWords() map[string]string {
	multipleWords, err := csvservice.ToMap("multiplewords.csv")
	if err != nil {
		return DefaultMultipleWords
	}

	if len(multipleWords) == 0 {
		return DefaultMultipleWords
	}

	return multipleWords
}

func GetSimilarWords() map[string]string {
	similarWords, err := csvservice.ToMap("similarwords.csv")
	if err != nil {
		return DefaultSimilarWords
	}

	if len(similarWords) == 0 {
		return DefaultSimilarWords
	}

	return similarWords
}
