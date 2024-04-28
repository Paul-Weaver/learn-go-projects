package main

import (
	"encoding/csv"
	"os"
)

// LoadWords loads a map of words from a CSV file.
// It takes a filename as input and returns a map[string]string containing the loaded words.
// The CSV file should have two columns: the first column represents the word, and the second column represents its translation.
// The function skips the header row and returns an error if there is any issue with opening or reading the file.
func LoadWords(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	wordMap := make(map[string]string)
	for _, record := range records[1:] { // Skip header row
		wordMap[record[0]] = record[1]
	}
	return wordMap, nil
}
