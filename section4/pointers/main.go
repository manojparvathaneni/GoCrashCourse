package main

import "fmt"

func main() {
	myInt := 10
	fmt.Println(myInt)

	// Declare a pointer 'myPointer1' and assign it the address of 'myInt'.
	myPointer1 := &myInt
	// Dereference 'myPointer1' to change the value of 'myInt' to 15.
	*myPointer1 = 15

	// Print the address of 'myInt', the value at the address stored in 'myPointer1',
	// and the value of 'myPointer1' itself (which is the address of 'myInt').
	fmt.Println(&myInt, *myPointer1, myPointer1)

	// Call 'changeValuOfInt', passing the address of 'myInt' to modify its value.
	changeValuOfInt(&myInt)
	fmt.Println(myInt)
}

// Define 'changeValuOfInt' with a pointer to an integer as a parameter.
func changeValuOfInt(num *int) {
	// Dereference 'num' to change the value at the referenced address to 20.
	*num = 20
}
