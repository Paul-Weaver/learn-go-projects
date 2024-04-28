package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: go run main.go <part-of-speech>")
		return
	}
	partOfSpeech := os.Args[1]
	filename := "vocab/" + partOfSpeech + ".csv"

	words, err := LoadWords(filename)
	if err != nil {
		fmt.Println("Failed to load words:", err)
		return
	}

	Quiz(words)
}
