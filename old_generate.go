package main

/*
import "fmt"

func firstWord() string {
	max := 0
	firstWord := ""

	for word, wordStats := range stats {
		if wordStats.IsFirst > max {
			max = wordStats.IsFirst
			firstWord = word
		}
	}

	return firstWord
}

func predictNextWord(lastWord string) string {
	max := 0
	nextWord := ""

	for word, wordStats := range stats {
		for prevWord, occur := range wordStats.PreviousWords[1] {
			if lastWord == prevWord {
				if occur > max {
					max = occur
					nextWord = word
				}
			}
		}
	}

	return nextWord
}

func playGround() {

	// Most obvious sentence
	sentence := ""
	word := firstWord()
	sentence += word + " "
	for i := 0; i < 100; i++ {
		word = predictNextWord(word)
		sentence += word + " "
	}
	fmt.Println(sentence)

	// Try to complete
	sentence = "Humans"
	word = "Humans"
	for i := 0; i < 100; i++ {
		word = predictNextWord(word)
		sentence += word + " "
	}
	fmt.Println(sentence)

	// Try to complete
	sentence = "Cat"
	word = "Cat"
	for i := 0; i < 100; i++ {
		word = predictNextWord(word)
		sentence += word + " "
	}
	fmt.Println(sentence)

}
*/
