package main

import "fmt"

func main() {
	x := 100
	var y int
	var ip *int // ip is a pointer to int

	ip = &x // ip now points to x
	y = *ip // y is now 1

	fmt.Println("ip = ", ip)
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	// new() function creates a variable and returns a pointer to the variable
	// that variable is initialised to Zero of it's type

	ptr := new(int)
	fmt.Println("ptr = ", *ptr)
	ptr = ip
	fmt.Println("ptr = ", *ptr)
	*ptr = 3
	fmt.Println("ptr = ", *ptr)
}
