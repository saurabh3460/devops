package main

import (
	"fmt"
)

type LinesOfText [][]byte

func main() {
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	for i, v := range text {
		fmt.Println(i, v)
	}
}
