package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Strings are immutable, but modified strings are returned

	x := "Hi Hi Hi There"
	y := "Hi There"
	fmt.Println("x = ", x)

	// STRINGS MANIPULATION

	comp := strings.Compare(x, y) // 0 if x==y; -1 if x<y; +1 if x>y
	fmt.Println(comp)

	index := strings.Index(y, "Th") // Finds first occurrence of substring in a string
	fmt.Println(index)

	newString := strings.Replace(x, "Hi", "Hello", 2) // Replaces n Instances of Old substring with New Substring in a String
	fmt.Println(newString)

	fmt.Println(strings.ToLower(x))
	fmt.Println(strings.ToUpper(y))

	fmt.Println(strings.TrimSpace(newString)) // Removes all Trailing and Leading Spaces

	//	STRING CONVERSIONS
	strs := "123"
	num, _ := strconv.Atoi(strs)
	fmt.Println(num + 321)

	ints := 123
	strnum := strconv.Itoa(ints)
	fmt.Println(strnum + "4 Let's go to Dance Floor")

}
