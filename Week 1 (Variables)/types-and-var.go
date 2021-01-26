package main

import "fmt"

// types are Aliases to the type of Variables like int, float, strings etc.
type Celsius float64
type IDNum int64

func main() {
	var temp Celsius
	temp = 100.45
	fmt.Println("Temperature = ", temp)

	var costumerId IDNum
	costumerId = 69
	fmt.Println("Customer ID = ", costumerId)

	// If variable is uninitialised, it's value will be Zero of it's type.
	// For Ex. 0 for int, 0.0 for float, "" for string
	var emptyInt int
	fmt.Println("emptyInt = ", emptyInt)
	var emptyString string
	fmt.Println("emptyString = ", emptyString)
	var emptyBool bool
	fmt.Println("emptyBool = ", emptyBool)

	// The := Operator (could be only done inside a function)
	x := 0.1
	fmt.Println("x = ", x)
}
