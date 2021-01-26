package main

import (
	"fmt"
)

func main() {
	var x int
	_, _ = fmt.Scanf("%d", &x)

	// If Else
	if x > 5 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	// Switch Case
	switch x {
	case 1:
		fmt.Println("case 1")
		break
	case 2:
		fmt.Println("case 2")
		break
	default:
		fmt.Println("case ", x)
	}

	// Tagless Switch
	switch {
	case x > 1:
		fmt.Println("x is greater than 1")
		break
	case x < 1:
		fmt.Println("x is lesser than 1")
		break
	}

	// For loop
	for x = 0; x < 5; x += 1 {
		fmt.Println("x = ", x)
	}

}
