package main

import (
	"fmt"
)

func printUniqueValue(arr []int) {
	//Create a dictionary for values of each element
	dict := make(map[int]int)

	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)
}

// Reverse returns its argument string reversed rune-wise left to right.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(reverse("!selpmaxe oG ,olleH"))
	inputArray := []int{10, 20, 30, 56, 67, 90, 10, 20}
	printUniqueValue(inputArray)
}
