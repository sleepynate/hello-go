package main

import (
	"fmt"
	. "strings"
)

func main() {
	fmt.Println("Please enter a candidate string for finding ian")
	input := ""
	_, _ = fmt.Scan(&input)
	if hasIan(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func hasIan(input string) bool {
	return HasPrefix(input, "i") && HasSuffix(input, "n") && Contains(input, "a")
}
