package main

import (
	"eliza/doctor"
	"fmt"
)

func main() {
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
}
