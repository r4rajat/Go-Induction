/*
	make()

	2 Argument Version
		1) Type of Slice
		2) Length/Capacity (B)oth are Same in this case
	Ex: sli := make([]int, 10)

	3 Argument Version
		1) Type of Slice
		2) Length
		3) Capacity	(Size of Underlying Array gets bigger than Slice
*/
/*
	append()
	can append elements to a slice and can add more elements than it's capacity
*/

package main

import "fmt"

func main() {
	sli := make([]int, 0, 3) //length of sli is 0 capacity is 3
	sli = append(sli, 100)
	sli = append(sli, 200)
	sli = append(sli, 300)
	sli = append(sli, 400)
	sli = append(sli, 500)
	for index, value := range sli {
		fmt.Printf("Index %d, Value %d\n", index, value)
	}
}
