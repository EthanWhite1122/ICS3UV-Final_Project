/*
Author:Ethan White
Version: 1.0.0
Date: 14-1-2026
Fileoverview:This program is a trivia program. 
One player enters how many questions they want to quiz player 2 on then inputs the questions and correct answers. 
Then the second player takes the quiz and tries to answer all the questions correctly. 
Finally it will print a final score at the end. 
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// variables
	correct := 0
	index := 0

	fmt.Println("Welcome to Trivia!")
	fmt.Println("This is a 2 player game where you make a mini quiz for your friend and they answer.")
	fmt.Println("Be sure to press enter after typing every input.")

	// number of questions
	var numQ int
	for {
		fmt.Print("How many questions would you like to quiz your friend on? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			numQ = value
			break
		}
		fmt.Println("Please enter a valid number greater than 0.")
	}

	// pre-sized slices
	questions := make([]string, numQ)
	answers := make([]string, numQ)

	// input questions
	questionNum := 1
	for index < numQ {
		fmt.Print("Enter question ", questionNum, ": ")
		question, _ := reader.ReadString('\n')
		question = strings.ToLower(strings.TrimSpace(question))

		fmt.Print("Enter the correct answer: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))

		if question == "" || answer == "" {
			fmt.Println("Question and answer cannot be empty.")
		} else {
			questions[index] = question
			answers[index] = answer
			index = index + 1
			questionNum = questionNum + 1
		}
	}

	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	// play quiz
	playIndex := 0
	for playIndex < numQ {
		fmt.Print(questions[playIndex], "? ")
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))

		if userAnswer == answers[playIndex] {
			fmt.Println("Correct!")
			correct = correct + 1
		} else {
			fmt.Println("Incorrect. The correct answer was:", answers[playIndex])
		}

		playIndex = playIndex + 1
	}

	// final results
	fmt.Println("Quiz complete. You got a score of", correct, "out of", numQ)
	fmt.Println("Done.")
}

