package main

import (
	"fmt"
)

// basic types(numbers, strings, booleans)
var myInt int
var myUint uint

var myFloat float32
var myFloat64 float64

func main() {

	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	fmt.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Manoj"
	fmt.Println(myString)
	myString = "Cooper" //strings are immutable in Go- this creates a new string.

	myBool := true
	var myBool1 bool

	fmt.Println(myBool, myBool1)

}
