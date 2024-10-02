package wordlist

import (
	"whatwords/src/csvservice"
)

func GetCustomWords() []string {
	customWords, err := csvservice.ToSlice("customwords.csv")
	if err != nil {
		return []string{}
	}

	if len(customWords) == 0 {
		return []string{}
	}

	return customWords
}

func GetExcludedWords() []string {
	excludedWords, err := csvservice.ToSlice("excludedwords.csv")
	if err != nil {
		return []string{}
	}

	if len(excludedWords) == 0 {
		return []string{}
	}

	return excludedWords
}

func GetShortWordsToKeep() []string {
	shortWordsToKeep, err := csvservice.ToSlice("shortwordstokeep.csv")
	if err != nil {
		return []string{}
	}

	if len(shortWordsToKeep) == 0 {
		return []string{}
	}

	return shortWordsToKeep
}

func GetMultipleWords() map[string]string {
	multipleWords, err := csvservice.ToMap("multiplewords.csv")
	if err != nil {
		return map[string]string{}
	}

	if len(multipleWords) == 0 {
		return map[string]string{}
	}

	return multipleWords
}

func GetSimilarWords() map[string]string {
	similarWords, err := csvservice.ToMap("similarwords.csv")
	if err != nil {
		return map[string]string{}
	}

	if len(similarWords) == 0 {
		return map[string]string{}
	}

	return similarWords
}
