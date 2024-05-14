package main

import (
	"myapp/packageone"
)

var myVar = "This is a package level variable in main."

func main() {
	blockVar := "This is a block level variable in main."

	packageone.PrintMe(myVar, blockVar)

}
