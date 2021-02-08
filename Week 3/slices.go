/*
-> Slice is a "window" on an underlying array
-> Variable Size, upto size of whole array
-> Pointer indicates the start of the slice
-> Length is the number of elements in slice
-> Capacity is the maximum number of elements(From Start of slice to end of Array)
*/

package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "e", "f", "g", "h"}
	slice1 := arr[1:3]
	slice2 := arr[2:5]

	fmt.Println("Slice 1")
	for index, value := range slice1 {
		fmt.Printf("Index %d, Value %s\n", index, value)
	}
	fmt.Println("Length of Slice 1: ", len(slice1))
	fmt.Println("Capacity of Slice 1: ", cap(slice1))
	fmt.Println("Slice 2")
	for index, value := range slice2 {
		fmt.Printf("Index %d, Value %s\n", index, value)
	}
	fmt.Println("Length of Slice 2: ", len(slice2))
	fmt.Println("Capacity of Slice 2: ", cap(slice2))

	// Slice Literals (Length and Capacity are Same)
	slice := []int{1, 2, 3}
	fmt.Println("Slice")
	for index, value := range slice {
		fmt.Printf("Index %d, Value %d\n", index, value)
	}
	fmt.Println("Length of Slice: ", len(slice))
	fmt.Println("Capacity of Slice: ", cap(slice))

}
