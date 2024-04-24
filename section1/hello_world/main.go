package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")
	sayHelloWorld("Hello world, from a new function")

	//variables
	//version 1 of declaration
	var message1 string
	message1 = "Hello from message1"
	sayHelloWorld(message1)

	//version 2 of declaration
	message2 := "Hello from message2"
	sayHelloWorld(message2)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
