package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var sliceVal string
	slice := make([]int, 0, 1)
	for {
		_, _ = fmt.Scanf("%s", &sliceVal)
		if sliceVal == "X" {
			break
		}
		intVal, _ := strconv.Atoi(sliceVal)
		slice = append(slice, intVal)
		sort.Ints(slice)
		fmt.Println(slice)

	}
}
