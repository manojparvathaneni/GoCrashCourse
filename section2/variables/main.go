package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just Press ENTER when ready"

func main() {
	// Seed the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate random numbers
	firstNumber := random.Intn(8) + 2  // Range: 2-9
	secondNumber := random.Intn(8) + 2 // Range: 2-9
	subtract := random.Intn(8) + 2     // Range: 2-9

	reader := bufio.NewReader(os.Stdin)

	// Game instructions
	fmt.Println("Guess the number game.")
	fmt.Println("----------------------")
	fmt.Println()

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtract, prompt)
	reader.ReadString('\n')

	// Calculate the predicted answer based on the sequence of operations
	answer := firstNumber*secondNumber - subtract
	fmt.Println("The answer is", answer)
}
