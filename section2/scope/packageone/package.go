package packageone

import "fmt"

var PackageVar = "I'm PackageVar Package variable in packageone "

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
