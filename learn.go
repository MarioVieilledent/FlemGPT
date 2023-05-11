package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const depth int = 1 // We take in account words until n-depth before the word analyzed, minimum is of course 1

var tokenArray []string = []string{}
var stats map[string]map[string][]int = map[string]map[string][]int{} // Unique instance of stats
var jsonOutput []byte = []byte{}                                      // Output to write in json

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

	// Create the array associating int to tokens
	fmt.Println("Taking in account all words...")
	createTokenArray(tokenizedInput)

	// Create the array of integers describing statistics about the words
	fmt.Println("Training...")
	train(tokenizedInput)

	// Compute the result (transform integers into float statistics)
	// fmt.Println("Conversion to JSON...")
	// convertToJSON()
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

// Create the array of all tokens encountered
func createTokenArray(tokenizedInput [][]string) {
	// Creation of the list of all word
	for _, paragraph := range tokenizedInput {
		for _, word := range paragraph {
			if word != "" && word != "\n" && word != "\r" {
				if !contains(tokenArray, word) {
					tokenArray = append(tokenArray, word)
				}
			}
		}
	}
	// Creation of stat array
	for _, word1 := range tokenArray {
		stats[word1] = map[string][]int{}
		for _, word2 := range tokenArray {
			stats[word1][word2] = make([]int, depth)
		}
	}
}

func train(tokenizedInput [][]string) {
	for _, paragraph := range tokenizedInput {
		for index, wordToAnalyze := range paragraph {
			if wordToAnalyze != "" && wordToAnalyze != "\n" && wordToAnalyze != "\r" {
				j := 0
				for i := index - 1; i >= index-depth; i-- {
					if i >= 0 {
						prevWord := paragraph[i]
						if stats[wordToAnalyze][prevWord] != nil {
							stats[wordToAnalyze][prevWord][j]++
						} else {
							fmt.Println("Nil []int for word: " + wordToAnalyze + " and prev: " + prevWord)
						}
					}
				}
				j++
			}
		}
	}
}

func convertToJSON() {
	res, err := json.Marshal(stats)
	if err != nil {
		panic("Error Marshalling")
	}
	jsonOutput = res
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
/*
func indexOf(arr []string, str string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return i
		}
	}
	return -1
}
*/
