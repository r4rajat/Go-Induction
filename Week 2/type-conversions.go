package main

import "fmt"

func main() {
	var x int32 = 1
	var y int16 = 2
	//	x = y this will give us an error because we can't do this explicitly
	x = int32(y)
	fmt.Println("x = ", x)

	var a float64 = 123.456
	var b float64 = 1.23456e2
	fmt.Println(a, b)

	var comp complex128 = complex(2, 3)
	fmt.Println(comp)
}
