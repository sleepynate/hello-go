package main

import "fmt"

func main() {

	fmt.Println("Enter a floating point number to truncate")
	var input float32
	var leng, err = fmt.Scan(&input)
	if leng == 0 || nil != err {
		fmt.Println("Sorry that won't work")
	} else {
		var asInt = int32(input)

		fmt.Printf("Truncated that is %d\n", asInt)
	}
}
