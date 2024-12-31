package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Text    string
	Options [4]string
	Answer  int
}

func takeQuiz(questions []Question, timeLimitPerQuestion time.Duration) int {
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for i, question := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Text)
		for j, option := range question.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}
		fmt.Println("Enter your answer (or type 'exit' to quit):")

		answerChan := make(chan string)
		go func() {
			input, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(input)
		}()

		timer := time.NewTimer(timeLimitPerQuestion)
		select {
		case <-timer.C:
			fmt.Println("Time's up for this question!")
			continue
		case answer := <-answerChan:
			if strings.ToLower(answer) == "exit" {
				fmt.Println("Exiting the quiz...")
				return score
			}

			option, err := strconv.Atoi(answer)
			if err != nil || option < 1 || option > 4 {
				fmt.Println("Invalid input. Please enter a number between 1 and 4.")
				continue
			}

			if option == question.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong answer.")
			}
		}
	}

	return score
}

func classifyPerformance(score, totalQuestions int) string {
	percentage := float64(score) / float64(totalQuestions) * 100
	switch {
	case percentage >= 80:
		return "Excellent"
	case percentage >= 50:
		return "Good"
	default:
		return "Needs Improvement"
	}
}

func main() {
	questions := []Question{
		{
			Text:    "What is the capital of France?",
			Options: [4]string{"Berlin", "Madrid", "Paris", "Rome"},
			Answer:  3,
		},
		{
			Text:    "What is 2 + 2?",
			Options: [4]string{"3", "4", "5", "6"},
			Answer:  2,
		},
		{
			Text:    "Who wrote 'Hamlet'?",
			Options: [4]string{"Charles Dickens", "William Shakespeare", "Jane Austen", "Mark Twain"},
			Answer:  2,
		},
	}

	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("You will have 10 seconds to answer each question.")

	score := takeQuiz(questions, 10*time.Second)

	fmt.Printf("\nQuiz Complete! You scored %d out of %d.\n", score, len(questions))
	performance := classifyPerformance(score, len(questions))
	fmt.Printf("Your performance: %s\n", performance)
}
