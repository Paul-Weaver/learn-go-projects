package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

// Question interface that all question types must implement
type Question interface {
	Display() // Display the question to the user
	CheckAnswer(input string) bool // Check if the user's answer is correct
}

// DirectInputQuestion for questions requiring direct text input
type DirectInputQuestion struct {
	Word	    string
	Translation string
}

// Display implements the display method for direct input questions
func (q DirectInputQuestion) Display() {
	fmt.Printf("Translate to Swedish: '%s'\n", q.Word)
}

// MultipleChoiceQuestion for questions that provide multiple choices
type MultipleChoiceQuestion struct {
	Word          string
	Choices       []string
	CorrectAnswer string
}

// Display implements the display method for multiple choice questions
func (q MultipleChoiceQuestion) Display() {
	fmt.Printf("What is the Swedish word for '%s'?\n", q.Word)
	for index, choice := range q.Choices {
		fmt.Printf("%d: %s\n", index+1, choice)
	}
}

// CheckAnswer verifies if the selected choice is the correct answer
func (q MultipleChoiceQuestion) CheckAnswer(input string) bool {
	choiceNum, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return q.Choices[choiceNum-1] == q.CorrectAnswer
}
