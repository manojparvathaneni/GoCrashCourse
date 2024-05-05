package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// numOfArgs := len(os.Args)
	// if numOfArgs <= 1 {
	// 	fmt.Println("Usage: echo <string_to_echo>")
	// } else {
	// 	stringToEcho := ""
	// 	for i := 1; i < numOfArgs; i++ {
	// 		stringToEcho += os.Args[i] + " "
	// 	}
	// 	fmt.Println(stringToEcho)
	// }

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: echo <string_to_echo>")
		os.Exit(1)
	}

	echoString := strings.Join(args, " ")
	fmt.Println(echoString)

}
