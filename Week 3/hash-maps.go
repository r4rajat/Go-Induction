/*
A hash table is a data structure used in a lot of different languages,
and it's very useful data structure, it allows you fast access to large bodies of data.
Hash table generally contains a key value pairs,
so there are a lot of values inside this hash table but each value is associated with a unique key.
The unique key and value and the key has to be unique that's actually completely important.
So, a hash table is meant to store these key value pairs, and it is important that each key is unique.
So a hash function is defined and is used to take a key and compute a slot in the hash table to insert the value according to the key.
*/

package main

import "fmt"

func main() {
	var idMap map[string]int
	idMap = make(map[string]int)
	idMap["a"] = 1
	idMap["b"] = 2
	idMap["c"] = 3
	fmt.Println(idMap["b"])
	idMap["b"] = 200
	fmt.Println(idMap["b"])
	delete(idMap, "b")
	id, p := idMap["b"] // id contains value and p is bool (true if value is mapped)
	fmt.Println(id, p)
	fmt.Println("Length of idMap: ", len(idMap)) //Length of Map

	for key, value := range idMap {
		fmt.Println(key, value)
	}

}
