package main

import "fmt"

func main() {
	var nop interface{}
	nop = int64(23)
	vl, ok := nop.(int64)
	fmt.Println(vl, ok)
}
