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

var ShortWordsToKeep = []string{
	"c",
}

var MultipleWords = map[string]string{
	"c plus plus": "cpp",
	"full stack":  "fullstack",
	"front end":   "frontend",
	"back end":    "backend",
	"web site":    "website",
	"web app":     "webapp",
	"web page":    "webpage",
	"web design":  "webdesign",
	"vue js":      "vuejs",
}

var SimilarWords = map[string]string{
	"c++": "cpp",
	"c#":  "csharp",
	"js":  "javascript",
	"ts":  "typescript",
	"py":  "python",
	"rb":  "ruby",
}

var ExtludedWords = []string{
	"the",
	"and",
	"for",
	"but",
	"nor",
	"les",
	"des",
	"une",
	"sur",
	"aux",
	"avec",
	"par",
	"pour",
	"chez",
	"vers",
	"depuis",
	"pendant",
	"avant",
	"apres",
	"pendant",
	"contre",
	"dans",
	"envers",
	"entre",
	"jusque",
	"sauf",
	"sous",
	"qui",
	"que",
	"quoi",
	"notre",
	"votre",
	"leur",
	"mon",
	"leurs",
	"vos",
	"nos",
	"son",
	"ses",
	"mes",
	"tous",
	"toutes",
	"tout",
	"toute",
	"ton",
	"nous",
	"vous",
	"ils",
	"elles",
	"elle",
	"est",
}

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

	wordparser.RemoveLineBreaks(&wordList)
	wordparser.ReplaceSimilarWords(&wordList, SimilarWords)
	wordparser.RemoveSpecialCharacters(&wordList)
	wordparser.MakeLowerCase(&wordList)
	wordparser.ReplaceMultipleWords(&wordList, MultipleWords)
	wordparser.RemoveWordShorterThanExcept(&wordList, 2, ShortWordsToKeep)
	wordparser.RemoveExcludedWords(&wordList, ExtludedWords)

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
