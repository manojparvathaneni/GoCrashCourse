package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs.\n", a.Name, a.NumberOfLegs)
}

func main() {

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()

	myTotal := sumMany(2, 3, 4, 5, 7, 8, 8)
	fmt.Println(myTotal)
}

// variadic functions.. the variadic parameter has to be at the end.
func sumMany(nums ...int) int {
	total := 0

	for _, num := range nums {

		total = total + num
	}
	return total
}
