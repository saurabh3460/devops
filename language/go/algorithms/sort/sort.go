package sort

import (
	"fmt"
	"log"
)

func init() {
	fmt.Println("Starting sorting")
}

func Sort(input []int) []int {
	inputLen := len(input)
	for i := 0; i < inputLen; i++ {
		for j := i + 1; j < inputLen; j++ {
			log.Printf("index: %d %d, values: %d %d", i, j, input[i], input[j])
			if input[i] > input[j] {
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
			}
		}
	}
	return input
}

/*
Need to fix

for i, iv := range input {
		for j := i + 1; j < inputLen; j++ {
			log.Printf("index: %d %d, values: %d %d", i, j, iv, input[j])
			if iv > input[j] {
				temp := iv
				input[i] = input[j]
				input[j] = temp
			}
		}
	}

output:
2021/04/23 22:34:05 index: 0 1, values: 2 4
2021/04/23 22:34:05 index: 0 2, values: 2 6
2021/04/23 22:34:05 index: 0 3, values: 2 3
2021/04/23 22:34:05 index: 0 4, values: 2 5
2021/04/23 22:34:05 index: 0 5, values: 2 74
2021/04/23 22:34:05 index: 1 2, values: 4 6
2021/04/23 22:34:05 index: 1 3, values: 4 3
2021/04/23 22:34:05 index: 1 4, values: 4 5
2021/04/23 22:34:05 index: 1 5, values: 4 74
2021/04/23 22:34:05 index: 2 3, values: 6 4
2021/04/23 22:34:05 index: 2 4, values: 6 5
2021/04/23 22:34:05 index: 2 5, values: 6 74
2021/04/23 22:34:05 index: 3 4, values: 6 6
2021/04/23 22:34:05 index: 3 5, values: 6 74
2021/04/23 22:34:05 index: 4 5, values: 6 74
[2 3 5 6 6 74]





*/
