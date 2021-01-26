package main

import "fmt"

func main() {
	const x = 1.34
	const ( //Multiple Initialization
		y = 4
		z = "Hello"
	)

	fmt.Println(x, y, z)
}
