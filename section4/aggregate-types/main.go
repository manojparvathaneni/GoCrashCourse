package main

import "fmt"

//aggregate types(array, struct)

// Structs
type Pet struct {
	Name  string
	Age   int
	Type  string
	Color string
	Breed string
}

func main() {

	// Arrays
	var myPack [4]string

	myPack[0] = "Cooper"
	myPack[1] = "Udo"
	myPack[2] = "Sydney"
	myPack[3] = "Azalea"

	// printing array using range
	for i, packMember := range myPack {
		fmt.Println(i, packMember)
	}

	//printing array using range, but ignoring index using the _
	for _, packMember := range myPack {
		fmt.Println(packMember)
	}

	//printing array using index
	for i := 0; i < len(myPack); i++ {
		fmt.Println(myPack[i])
	}

	//structs
	var cooper Pet
	cooper.Name = "Cooper"
	cooper.Age = 9
	cooper.Type = "Dog"
	cooper.Color = "Brown"
	cooper.Breed = "GSD mix"

	udo := Pet{
		Name:  "Udo",
		Age:   8,
		Type:  "Dog",
		Color: "Black",
		Breed: "Pointer/Retreiver mix",
	}

	fmt.Println(cooper, udo)
}
