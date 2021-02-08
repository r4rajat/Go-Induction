package main

import "fmt"

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {
	var p1 = Person{
		"Rajat Gupta",
		"Rajasthan-301404",
		"7508415858",
	}

	var p2 Person
	p2.name = "Aishwarya"
	p2.addr = "Pune"
	p2.phone = "6464679565"

	fmt.Println(p1.name)
	fmt.Println(p2.addr)
}
