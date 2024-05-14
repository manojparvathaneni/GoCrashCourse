package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Capuccino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Machiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte ")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Machiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		if i >= 1 && i <= 6 {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		} else {
			fmt.Println("Invalid Menu choice")
		}

	}

	fmt.Println("Program Exiting..")
}
