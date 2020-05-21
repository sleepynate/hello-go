package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	marshalMap := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a name")
	if scanner.Scan() {
		marshalMap["name"] = scanner.Text()
	}
	fmt.Println("... and now an address would be swell")
	if scanner.Scan() {
		marshalMap["address"] = scanner.Text()
	}

	jsonPerson, _ := json.Marshal(marshalMap)

	fmt.Println("Great, here is some useless JSON on your terminal instead of in a browser")
	fmt.Println(string(jsonPerson))
}
