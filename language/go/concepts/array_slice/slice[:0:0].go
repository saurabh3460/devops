package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	d := []string{"r", "o", "a", "d", "j", "o", "r"}
	e := d[1:4:4]
	// e := d[1:4]
	fmt.Println(e, len(e), cap(e))

	var dstCap [3]int
	copiedSLice := dstCap[:0] // creating empty slice using [:0] with cap 3
	// copiedSLice[0] = 3
	fmt.Println(copiedSLice, len(copiedSLice), cap(copiedSLice), copiedSLice == nil)

	cleanCopy := dstCap[:0:0]
	fmt.Println(cleanCopy, len(cleanCopy), cap(cleanCopy))

	// fmt.Println(&copiedSLice[0], &copiedSLice[0])

	var DstCap [2]int
	fmt.Println(DstCap)

}

/*

Find out other way to use [:0] and [:0:0] to create slice. Use this reference for use case:
https://cs.opensource.google/go/go/+/refs/tags/go1.18.3:src/regexp/regexp.go;l=821
https://cs.opensource.google/go/go/+/refs/tags/go1.18.3:src/regexp/exec.go;drc=c1a4e0fe014568501b194eb8b04309f54eee6b4c;l=521



*/
