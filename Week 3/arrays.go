package main

import "fmt"

func main() {
	var arr1 [5]int
	arr1[0] = 2
	fmt.Println(arr1[0]) //Initialized 2
	fmt.Println(arr1[1]) // Not Initialized so Default will be 0

	//	Array Literals

	y := [...]int{1, 2, 3, 4, 5} // ... for size in Array Literals infers size from number of initializers

	// Iterating Through Arrays

	for index, value := range y {
		fmt.Printf("Index %d, Value %d\n", index, value)
	}

}
