package main

import "fmt"

func main() {
	var floatVariable float64
	_, _ = fmt.Scanf("%f", &floatVariable)

	var truncateInt int64

	truncateInt = int64(floatVariable)
	fmt.Println("Truncated Integer : ", truncateInt)

}
