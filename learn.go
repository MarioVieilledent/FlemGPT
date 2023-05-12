package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const depth int = 8 // We take in account words until n-depth before the word analyzed, minimum is of course 1

var tokenArray []string = []string{}
var stats map[string]map[string][]int = map[string]map[string][]int{} // Unique instance of stats
var jsonOutput []byte = []byte{}                                      // Output to write in json

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

		// Reduce as much as possible number of tokens (by removing some chars, putting everything in lowerCase)
		for k, word := range words {
			words[k] = simplify(word)
		}
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

	fmt.Println("Number of words: " + strconv.Itoa(len(tokenArray)))

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
						if prevWord != "" && prevWord != "\n" && prevWord != "\r" {
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

func simplify(word string) string {
	simplified := strings.ToLower(word)
	simplified = strings.ReplaceAll(simplified, "\n", "")
	simplified = strings.ReplaceAll(simplified, "\r", "")
	simplified = strings.ReplaceAll(simplified, ".", "")
	simplified = strings.ReplaceAll(simplified, ",", "")
	simplified = strings.ReplaceAll(simplified, ";", "")
	simplified = strings.ReplaceAll(simplified, "?", "")
	simplified = strings.ReplaceAll(simplified, "!", "")
	simplified = strings.ReplaceAll(simplified, ":", "")
	simplified = strings.ReplaceAll(simplified, "'", "")
	simplified = strings.ReplaceAll(simplified, "-", "")
	simplified = strings.ReplaceAll(simplified, "_", "")
	simplified = strings.ReplaceAll(simplified, "(", "")
	simplified = strings.ReplaceAll(simplified, ")", "")
	simplified = strings.ReplaceAll(simplified, "[", "")
	simplified = strings.ReplaceAll(simplified, "]", "")
	simplified = strings.ReplaceAll(simplified, "{", "")
	simplified = strings.ReplaceAll(simplified, "}", "")
	simplified = strings.ReplaceAll(simplified, "~", "")
	simplified = strings.ReplaceAll(simplified, "&", "")
	simplified = strings.ReplaceAll(simplified, "#", "")
	simplified = strings.ReplaceAll(simplified, "\"", "")
	simplified = strings.ReplaceAll(simplified, "|", "")
	simplified = strings.ReplaceAll(simplified, "^", "")
	simplified = strings.ReplaceAll(simplified, "@", "")
	simplified = strings.ReplaceAll(simplified, "=", "")
	simplified = strings.ReplaceAll(simplified, "+", "")
	simplified = strings.ReplaceAll(simplified, "/", "")
	simplified = strings.ReplaceAll(simplified, "*", "")
	simplified = strings.ReplaceAll(simplified, "%", "")
	simplified = strings.ReplaceAll(simplified, "\\", "")
	simplified = strings.ReplaceAll(simplified, "0", "")
	simplified = strings.ReplaceAll(simplified, "1", "")
	simplified = strings.ReplaceAll(simplified, "2", "")
	simplified = strings.ReplaceAll(simplified, "3", "")
	simplified = strings.ReplaceAll(simplified, "4", "")
	simplified = strings.ReplaceAll(simplified, "5", "")
	simplified = strings.ReplaceAll(simplified, "6", "")
	simplified = strings.ReplaceAll(simplified, "7", "")
	simplified = strings.ReplaceAll(simplified, "8", "")
	simplified = strings.ReplaceAll(simplified, "9", "")
	return simplified
}
