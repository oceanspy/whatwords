package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"whatwords/src/color"
	"whatwords/src/message"
	"whatwords/src/wordlist"
	"whatwords/src/wordparser"
)

func main() {
	// if os.Stdin is empty, print a message and exit
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		message.Error("No input detected. Stopping.")
		message.Info("Did you correctly pipe whatwords into a valid text source ?")
		message.Text("Example: cat myFile.txt | whatwords")
		os.Exit(0)
	}

	maxWordsToShow := 20
	if len(os.Args) == 2 {
		if _, err := strconv.Atoi(os.Args[1]); err == nil {
			maxWordsToShow, _ = strconv.Atoi(os.Args[1])
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

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

	wordparser.MakeLowerCase(&wordList)
	wordparser.SplitWordsByDelimiters(&wordList)
	wordparser.RemoveLineBreaks(&wordList)
	wordparser.RemoveEmptyWords(&wordList)
	wordparser.ReplaceSimilarWords(&wordList, wordlist.GetSimilarWords())
	wordparser.RemoveSpecialCharacters(&wordList)
	wordparser.ReplaceMultipleWords(&wordList, wordlist.GetMultipleWords())
	wordparser.RemoveWordShorterThanExcept(&wordList, 2, wordlist.GetShortWordsToKeep())
	wordparser.RemoveExcludedWords(&wordList, wordlist.GetExcludedWords())

	wordsWithInfos := wordparser.CalculateOccurenceOfEachWordInsideSlice(&wordList)
	wordparser.SortByCount(&wordsWithInfos)

	// Print the result
	PrintTitle("Most used words")
	for i, e := range wordsWithInfos {
		if i > maxWordsToShow {
			break
		}

		PrintRow(e)
	}

	// Get the usage of predefined list
	customWords := wordlist.GetCustomWords()
	if len(customWords) > 0 {
		customWordsWithInfos := wordparser.CalculateOccurenceOf(&customWords, &wordList)
		wordparser.SortByCount(&customWordsWithInfos)

		PrintTitle("Custom List of words")
		for _, e := range customWordsWithInfos {
			PrintRow(e)
		}
	}
}

func PrintTitle(title string) {
	message.Ln()
	message.FixedTextLength(10, ' ', " COUNT")
	message.Text(strings.ToUpper(title))
	message.LineOf('─')
}

func PrintRow(e wordparser.WordInfo) {
	wordToPrint := strings.ReplaceAll(e.Word, "+", " ")
	if e.Count == 0 {
		fmt.Print(color.Red)
		message.FixedTextLength(10, ' ', " ", strconv.Itoa(e.Count))
		fmt.Print(color.Reset)
		message.Text(wordToPrint)
	} else if e.Count <= 1 {
		fmt.Print(color.Magenta)
		message.FixedTextLength(10, ' ', " ", strconv.Itoa(e.Count))
		fmt.Print(color.Reset)
		message.Text(wordToPrint)
	} else if e.Count <= 5 {
		fmt.Print(color.Yellow)
		message.FixedTextLength(10, ' ', " ", strconv.Itoa(e.Count))
		fmt.Print(color.Reset)
		message.Text(wordToPrint)
	} else if e.Count <= 10 {
		fmt.Print(color.Cyan)
		message.FixedTextLength(10, ' ', " ", strconv.Itoa(e.Count))
		fmt.Print(color.Reset)
		message.Text(wordToPrint)
	} else if e.Count > 10 {
		fmt.Print(color.Green)
		message.FixedTextLength(10, ' ', " ", strconv.Itoa(e.Count))
		fmt.Print(color.Reset)
		message.Text(wordToPrint)
	} else {
		fmt.Print(color.Gray)
		message.FixedTextLength(10, ' ', " ", strconv.Itoa(e.Count))
		fmt.Print(color.Reset)
		message.Text(wordToPrint)
	}
}
