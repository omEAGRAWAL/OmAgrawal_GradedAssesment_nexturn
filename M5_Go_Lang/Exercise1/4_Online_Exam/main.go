package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Question struct represents a quiz question
type Question struct {
	Question string
	Options  []string
	Answer   int // Index of the correct answer (1-based)
}

var questionBank = []Question{
	{"What is the capital of France?", []string{"1. Berlin", "2. Madrid", "3. Paris", "4. Rome"}, 3},
	{"What is 2 + 2?", []string{"1. 3", "2. 4", "3. 5", "4. 6"}, 2},
	{"Who wrote 'Hamlet'?", []string{"1. Charles Dickens", "2. William Shakespeare", "3. J.K. Rowling", "4. Mark Twain"}, 2},
}

const questionTimeLimit = 10 // Time limit for each question in seconds

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("You can type 'exit' anytime to quit the quiz.")

	score := 0
	totalQuestions := len(questionBank)

	for i, question := range questionBank {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Question)
		for _, option := range question.Options {
			fmt.Println(option)
		}

		fmt.Printf("Enter your choice (1-%d): ", len(question.Options))

		answerChan := make(chan int, 1)
		errorChan := make(chan error, 1)

		go func() {
			scanner.Scan()
			input := scanner.Text()

			// Check for "exit" command
			if strings.ToLower(input) == "exit" {
				errorChan <- fmt.Errorf("quiz exited")
				return
			}

			// Parse user input
			answer, err := strconv.Atoi(input)
			if err != nil || answer < 1 || answer > len(question.Options) {
				errorChan <- fmt.Errorf("invalid input")
				return
			}

			answerChan <- answer
		}()

		select {
		case <-time.After(time.Duration(questionTimeLimit) * time.Second):
			fmt.Println("\nTime's up for this question!")
			continue
		case err := <-errorChan:
			if err.Error() == "quiz exited" {
				fmt.Println("\nExiting the quiz...")
				goto end
			}
			fmt.Println("Invalid input. Skipping to the next question...")
			continue
		case answer := <-answerChan:
			if answer == question.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong answer!")
			}
		}
	}

end:
	// Score Calculation
	fmt.Println("\nQuiz Finished!")
	fmt.Printf("Your Score: %d/%d\n", score, totalQuestions)

	// Performance Classification
	if score == totalQuestions {
		fmt.Println("Performance: Excellent!")
	} else if score > totalQuestions/2 {
		fmt.Println("Performance: Good!")
	} else {
		fmt.Println("Performance: Needs Improvement!")
	}

	fmt.Println("Thank you for taking the quiz!")
}
