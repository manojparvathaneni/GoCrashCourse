package main

import (
	"fmt"
	"log"

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

	char, _, _ := keyboard.GetSingleKey()
	//i , _ := strconv.Atoi(string(char))
	fmt.Println(string(char))

}
