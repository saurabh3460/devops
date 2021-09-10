package main

import (
	"fmt"
	"runtime"
)

// https://github.com/hashicorp/go-version

func main() {
	fmt.Println(runtime.NumCPU())
}
