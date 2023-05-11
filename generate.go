package main

import "fmt"

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
	for i := len(previousWords) - 1; i >= len(previousWords)-depth; i-- {
		for word, stats := range wordStats {
			if previousWords[i] == word {
				score += 1.0 * (1.0 / float64(len(previousWords)-i))
				if len(stats) > len(previousWords)-depth-1 {
					score += float64(stats[len(previousWords)-depth-1])
				}
			}
		}
	}
	return score
}

func playGround() {
	var word string
	sentence := []string{"a", "programming", "language", "is"}
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
	word = findNexWord(sentence)
	sentence = append(sentence, word)
	fmt.Println(sentence)
}
