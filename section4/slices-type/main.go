package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {

	var fabfour []string
	fabfour = append(fabfour, "Cooper")
	fabfour = append(fabfour, "Udo")
	fabfour = append(fabfour, "Sydney")
	fabfour = append(fabfour, "Azalea")

	fmt.Println(fabfour)
	for _, v := range fabfour {
		fmt.Println(v)
	}
	fmt.Println(fabfour[3])
	fmt.Println(fabfour[0:2])
	fmt.Println(len(fabfour))
	fmt.Println(sort.StringsAreSorted(fabfour))
	sort.Strings(fabfour)
	fmt.Println(fabfour)
	fabfour = slices.Delete(fabfour, 0, 3)
	fmt.Println(fabfour)
}

func DeleteFromSlice(f []string, i int) []string {
	f[i] = f[len(f)-1]
	f[len(f)-1] = ""
	f = f[:len(f)-1]
	return f
}
