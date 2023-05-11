package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findNexWord(previousWords []string) string {
	bestFittingWord := ""
	bestScore := 0.0
	score := 0.0
	for word, wordStats := range stats {
		score = calculateScore(previousWords, wordStats)
		if score > bestScore {
			bestScore = score
			bestFittingWord = word
		}
	}
	return bestFittingWord
}

func calculateScore(previousWords []string, wordStats map[string][]int) float64 {
	score := 0.0
	for i := len(previousWords) - 1; i >= len(previousWords)-depth && i >= 0; i-- {
		for word, stats := range wordStats {
			if previousWords[i] == word {
				score += 1.0 * (1.0 / float64(len(previousWords)-i))
				if len(stats) > len(previousWords)-depth-1 && len(previousWords)-depth-1 > 0 {
					score += float64(stats[len(previousWords)-depth-1])
				}
			}
		}
	}
	return score
}

func loop() {
	reader := bufio.NewReader(os.Stdin)
	nbWords := 15
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
