package main

/*
import (
	"encoding/json"
	"fmt"
	"strings"
)

// Paragraph = A whole bunch of text analyzed, can be one or multiple paragraph, it is like all the answer of ChatGPT to a prompt

const depth int = 1 // We take in account words until n-depth before the word analyzed, minimum is of course 1

var stats rawInts = rawInts{} // Unique instance of stats
var jsonOutput []byte = []byte{}

type rawInts = map[string]wordStats

type wordStats struct {
	IsFirst       int              `json:"isFirst"`       // Number of times the word appears at the beginning of a paragraph
	IsLast        int              `json:"isLast"`        // Number of times the word appears at the end of a paragraph
	PreviousWords []map[string]int `json:"previousWords"` // List (n-1, n-2) of map of all words possibles and their occurrences before the word analyzed
}

func learn(input string) {
	// Tokenise input
	fmt.Println("Tokenization...")
	tokenizedInput := tokenize(input)

	// Create the array of integers describing statistics about the words
	fmt.Println("Training...")
	train(tokenizedInput)

	// Compute the result (transform integers into float statistics)
	fmt.Println("Computing...")
	compute()
}

// Accepts any UTF-8 file, returns a succession of token (similar to words)
func tokenize(input string) [][]string {
	// First list a a list of paragraph, or line
	// Second list is a list of token (just a string)

	tokens := [][]string{}

	lines := strings.Split(strings.ToLower(input), "\n")

	for k, line := range lines {
		tokens = append(tokens, []string{})

		words := strings.Split(line, " ")
		tokens[k] = words
	}

	return tokens
}

func compute() {
	res, err := json.Marshal(stats)
	if err != nil {
		panic("Error Marshalling")
	}
	jsonOutput = res
}

func train(tokenizedInput [][]string) {
	for _, paragraph := range tokenizedInput {
		for k, word := range paragraph {

			if word != "" && word != "\n" && word != "\r" {

				if w, ok := stats[word]; ok {
					if k == 0 {
						w.IsFirst++
						stats[word] = w
					}
					if k == len(paragraph)-1 {
						w.IsLast++
						stats[word] = w
					}
					incrementStats(paragraph, k, word)
				} else {
					previousWords := newWord(paragraph, k, word)

					isFirst := 0
					isLast := 0
					if k == 0 {
						isFirst = 1
					}
					if k == len(paragraph)-1 {
						isLast = 1
					}

					stats[word] = wordStats{
						IsFirst:       isFirst,
						IsLast:        isLast,
						PreviousWords: previousWords,
					}

					incrementStats(paragraph, k, word)
				}

			}
		}
	}
}

func newWord(paragraph []string, k int, word string) []map[string]int {
	previousWords := []map[string]int{}

	for i := 0; i <= depth; i++ { // Depth of 4, append 5 elem to array, because index 0 is unused, index 1 refers to word n-1, etc.
		previousWords = append(previousWords, map[string]int{})
	}

	return previousWords
}

func incrementStats(paragraph []string, k int, word string) {
	// Analyses the n-1 word
	if k >= 1 {
		prev := paragraph[k-1]

		if ws, ok := stats[word]; ok {
			if _, exists := ws.PreviousWords[1][prev]; exists {
				ws.PreviousWords[1][prev]++
			} else {
				ws.PreviousWords[1][prev] = 1
			}
		}
	}
}

// Checks if array contain a string
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}

	return false
}

// Get index of elem in array
func indexOf(arr []string, str string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return i
		}
	}
	return -1
}

*/
