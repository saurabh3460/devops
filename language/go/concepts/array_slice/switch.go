package main

import (
	"fmt"
)

func main() {

	name := "knell"

	switch name {
	case "bell":
		fmt.Println("moo")
	case "tacoo":
		fmt.Println("sue")
	case "toll", "chime", "knell":
		fmt.Println("Taco Bell itself")
	}

}
