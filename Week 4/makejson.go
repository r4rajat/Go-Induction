package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func userInput(prompt string) string {
	fmt.Println(prompt)

	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}
	return strings.TrimSpace(userInput)
}

func main() {
	nameAddrMap := map[string]string{}

	nameAddrMap["name"] = userInput("Enter Name:")
	nameAddrMap["address"] = userInput("Enter Address:")

	bytes, err := json.Marshal(nameAddrMap)
	if err == nil {
		fmt.Println(string(bytes))
	}

}
