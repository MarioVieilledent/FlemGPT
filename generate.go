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
			score = calculateScore(previousWords, wordStats)
			wordsScores = append(wordsScores, WordScore{word, score})
		}
	}

	sort.Slice(wordsScores, func(i, j int) bool {
		return wordsScores[i].score > wordsScores[j].score
	})

	return wordsScores[rand.Intn(3)].word
}

func calculateScore(previousWords []string, wordStats map[string][]int) float64 {
	score := 0.0
	j := 0
	for i := len(previousWords) - 1; i >= len(previousWords)-depth; i-- {
		if i > 0 {
			for word, stats := range wordStats {
				if word == previousWords[i] {
					if j == 0 {
						score += 15.0 + float64(stats[j])*0.3
					}
					if j == 1 {
						score += 3.0 + float64(stats[j])*0.1
					}
					if j == 2 {
						score += 1.0 + float64(stats[j])*0.05
					}
					score += 2.0*(float64(stats[j])*0.9) + (float64(depth-j) * float64(depth-j) * 0.5) + (float64(len(word)) * 1.5)
				}

				for k := len(previousWords) - 1; k >= len(previousWords)-depth; k-- {
					if word == previousWords[i] && k >= 0 {
						score += 0.2*float64(stats[k])*0.1 +
							float64(depth-k)*float64(depth-k)*0.4 +
							float64(len(word))*0.5
					}
				}
			}
		}
		j++
	}
	return score
}

func loop() {
	reader := bufio.NewReader(os.Stdin)
	nbWords := 8
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

func continueText(sentence []string, nbWords int) {
	var word string

	for _, w := range sentence {
		fmt.Print(w + " ")
	}

	for i := 0; i < nbWords; i++ {
		word = findNexWord(sentence)
		fmt.Print(word + " ")
		sentence = append(sentence, word)
	}

	fmt.Println("\n\n result:")
	fmt.Println(sentence)
}
