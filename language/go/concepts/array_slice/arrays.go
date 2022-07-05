package main

import (
	"fmt"
)

func main() {

	anArray := [4]string{}
	fmt.Println(len(anArray))

	// 2d array
	_2d := [3][2]int{{32, 43}, {75, 23}, {50, 89}}

	// fmt.Println(len(_2d))
	fmt.Println("------2D array------")

	// typical way to iterate an array
	for i := 0; i < len(_2d); i++ {
		for j := 0; j < len(_2d[i]); j++ {
			fmt.Println(_2d[i][j])
		}
		// fmt.Println(_2d[i])
	}

	fmt.Println("------Idiomatic way to interate an array------")

	// idiomatic way to interate an array
	for _, ivalue := range _2d {
		for _, jvalue := range ivalue {
			fmt.Println(jvalue)
		}
	}
}
