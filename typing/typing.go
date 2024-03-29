package typing

import (
	"fmt"
)

func ex_typing() {

	// total
	var totalScore int = 0

	// question array
	questions := []string{
		"hello",
		"bonjour",
		"こんにちは",
	}

	// ask questions
	for i := 0; i < len(questions); i++ {
		Ask(i+1, questions[i], &totalScore)
	}

	// show result
	fmt.Printf("Result: %d\n", totalScore)
}

func Ask(quizNo int, content string, scorePtr *int) {

	// user input
	var input string

	// show question and get user input
	fmt.Printf("Question %d: %s\n", quizNo, content)
	fmt.Scan(&input)

	// score
	if input == content {
		fmt.Printf("Correct!")
		*scorePtr += 10
	} else {
		fmt.Printf("Incorrect :(")
	}

	fmt.Printf("\n\n")

}
