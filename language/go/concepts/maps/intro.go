package main

import (
	"fmt"
	"strings"
)

func main() {
	var m map[string]int     // this will create an empty map and will not allow us to add items
	m = make(map[string]int) // this way we create a map
	fmt.Println(m)
	m["jake"] = 1
	m["simon"] = 2
	m["joe"] = 3
	fmt.Println(m)
	delete(m, "simon")
	fmt.Println(m)

	elem, ok := m["joe"]
	if !ok {
		fmt.Println("Not found", elem)
	} else {
		fmt.Println("Found", elem)
	}

	statement := "hey it is map usage"
	result := WordCount(statement)
	fmt.Println(result)
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, value := range words {
		counts[value] = 1
	}
	return counts
}
