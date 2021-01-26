/*
	iota is just like enumeration in C/C++
*/

package main

import "fmt"

type Grades int

func main() {

	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)
	fmt.Printf("A = %d, B = %d, C = %d, D = %d, E = %d, F = %d\n", A, B, C, D, E, F)
}
