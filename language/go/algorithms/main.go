package main

import (
	"fmt"

	"github.com/saurabh3460/algorithms/sort"
)

// func Reverse(input []int) []int {
// 	inputLen := len(input)
// 	output := make([]int, inputLen)

// 	for i, n := range input {
// 		j := inputLen - i - 1
// 		output[j] = n
// 	}
// 	return output
// }

func main() {
	var input = []int{2, 4, 6, 3, 5, 74}
	fmt.Println(sort.Sort(input))
}
