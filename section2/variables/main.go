package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and Press ENTER when ready"

func main() {

	firstNumber := 2
	secondNumber := 5
	subtract := 7
	var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the number game.")
	fmt.Println("----------------------")
	fmt.Println()

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now Multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtract, prompt)
	reader.ReadString('\n')

	answer = firstNumber*secondNumber - subtract
	fmt.Println("The answer is", answer)
}
