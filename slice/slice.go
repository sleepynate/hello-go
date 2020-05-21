package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	var s = make([]int, 0, 3)
	for {
		fmt.Println("Enter integers or X to quit")
		_, _ = fmt.Scanln(&input)
		if input == "X" { break }
		asInt, _ := strconv.Atoi(input)
		s = append(s, asInt)
		sort.Ints(s)
		for i, v := range s {
			fmt.Print(v)
			if i != len(s) - 1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("")
	}
}
