package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

type WordScore struct {
	word  string
	score float64
}

func findNexWord(previousWords []string) string {
	score := 0.0
	wordsScores := []WordScore{}
	for word, wordStats := range stats {
		if word != " " && word != "\n" && word != "\r" {
			score = calculateScore(word, previousWords, wordStats)
			wordsScores = append(wordsScores, WordScore{word, score})
		}
	}

	sort.Slice(wordsScores, func(i, j int) bool {
		return wordsScores[i].score > wordsScores[j].score
	})

	return wordsScores[rand.Intn(randPickOverBest)].word
}

func calculateScore(wordN string, previousWords []string, wordStats map[string][]int) float64 {
	score := 0.0
	nMinus := 0
	if val, ok := stats[wordN]; ok {
		if val2, ok := val[previousWords[len(previousWords)-1]]; ok {
			if val2[0] >= minOccurToAccept {
				score += 1.0
				for i := len(previousWords) - 1; i >= len(previousWords)-depth; i-- {
					if i > 0 {
						for word, stats := range wordStats {
							if word == previousWords[i] {
								score += (coefOccur * float64(stats[nMinus])) *
									(coefDistance * float64(depth-nMinus) * float64(depth-nMinus)) *
									(coefWordLength * float64(len(word)))
							}
						}
					}
					nMinus++
				}
			}
		}
	}
	return score
}

func loop() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Prompt: ")
		input, _ := reader.ReadString('\n')
		input = strings.ReplaceAll(input, "\n", "")
		input = strings.ReplaceAll(input, "\r", "")
		sentence := strings.Split(input, " ")
		for k, v := range sentence {
			sentence[k] = simplify(v)
		}
		continueText(sentence, nbWords)
	}
}

func continueText(sentence []string, nbWords int) []string {
	var word string

	for _, w := range sentence {
		fmt.Print(w + " ")
	}

	for i := 0; i < nbWords; i++ {
		word = findNexWord(sentence)
		fmt.Print(word + " ")
		sentence = append(sentence, word)
	}

	// fmt.Println("\n\n result:")
	// fmt.Println(sentence)

	return sentence
}
