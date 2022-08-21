package main

import (
	"errors"
	"fmt"
	"os"
)

func Load(filename string) error {
	file, error := os.ReadFile(filename)

	if error != nil {
		if errors.Is(error, os.ErrNotExist) {
			// fmt.Print("No file found")
			return error
		}
		return error
	}

	if len(file) == 0 {
		return errors.New("File is empty")
	}

	return nil

}

func main() {
	fmt.Println(Load("./test.txt"))
}
