package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/fatih/color"
)

// Quiz is a function that conducts a vocabulary quiz in Swedish.
// It takes a map of words as input, where the keys are the English words
// and the values are the corresponding Swedish words. The function randomly
// selects a word from the map and prompts the user to provide the Swedish
// translation. If the user's answer matches the correct translation, the
// function prints "Correct!" and increments the score. If the answer is
// incorrect, the function prints the correct translation. After all the
// words have been tested, the function prints the final score.
func Quiz(words map[string]string) {
	rand.Seed(time.Now().UnixNano())
	keys := make([]string, 0, len(words))
	for k  := range words {
		keys = append(keys, k)
	}

	score := 0
	for i := 0; i < len(keys); i++ {
		// Randomise index
		idx := rand.Intn(len(keys))
		key := keys[idx]

		fmt.Printf("What is the Swedish word for '%s'? ", key)
		var answer string
		fmt.Scanln(&answer)

		if answer == words[key] {
			color.Green("Correct!")
			score++
		} else {
			color.Red("Wrong! The correct answer is '%s'.\n", words[key])
		}
	}

	fmt.Printf("Quiz finished! Your score: %d/%d\n", score, len(words))
}
