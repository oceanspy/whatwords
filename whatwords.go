package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"whatwords/src/message"
	"whatwords/src/wordparser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(SplitByDelimiters)

	var wordList []string

	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		message.Error("Couldn't read the input. Stopping.")
		message.Info("Did you correctly pipe whatwords into a valid text source ?")
		message.Text("Example: cat myFile.txt | whatwords")
		message.Ln()
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(0)
	}

	wordparser.RemoveSpecialCharactersFromList(&wordList)
	wordparser.MakeLowerCaseFromList(&wordList)
	wordparser.RemoveWordShorterThanExcept(&wordList, 2)
	// wordparser.ReplaceSimilarWords(&wordList)

	wordsWithInfos := wordparser.CalculateOccurenceOfEachWordInsideSlice(&wordList)
	wordparser.SortByCount(&wordsWithInfos)

	// Get the usage of predefined list

	// Print the result
	message.Ln()
	message.FixedTextLength(10, ' ', " COUNT")
	message.Text("MOST USED WORDS")
	message.LineOf('─')
	for i, e := range wordsWithInfos {
		if i > 20 {
			break
		}

		message.FixedTextLength(10, ' ', " ", strconv.Itoa(e.Count))
		message.Text(e.Word)
	}
}

func SplitByDelimiters(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Use a regular expression to find words, considering spaces and commas as delimiters
	re := regexp.MustCompile(`[ ,;:.!?]+`) // Matches space, comma, semicolon, period, exclamation, and question mark

	// Split the data by the regular expression
	tokens := re.Split(string(data), -1)

	// Iterate through the tokens and return the first non-empty token
	for _, token := range tokens {
		if token != "" {
			advance = len(token)
			return advance, []byte(token), nil
		}
	}

	if atEOF {
		return 0, nil, nil // No more data and no matches
	}
	return 0, nil, nil // Not at EOF, and no matches
}
