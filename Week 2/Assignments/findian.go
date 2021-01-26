package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	_, _ = fmt.Scanln(&inputString)
	ifIPresent := strings.Index(inputString, "i")
	ifAPresent := strings.Index(inputString, "a")
	ifNPresent := strings.Index(inputString, "n")
	if ifAPresent != -1 && ifIPresent != -1 && ifNPresent != -1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
