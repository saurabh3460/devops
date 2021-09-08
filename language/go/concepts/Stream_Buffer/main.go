package main

import (
	"fmt"
	"io"
)

type MyStringData struct {
	str string
	readIndex int
}

func (msd *MyStringData) Read(p []byte) (n int, err error) {

	sourceBytes := []byte(msd.str)
	if msd.readIndex >= len(sourceBytes) {
		return 0, io.EOF
	}
	nextReadLimit := msd.readIndex + len(p)
	if nextReadLimit >= len(sourceBytes) {
		nextReadLimit = len(sourceBytes)
		err = io.EOF
	}
	bytes := sourceBytes[msd.readIndex:nextReadLimit]
	n = len(bytes)
	copy(p, bytes)
	msd.readIndex = nextReadLimit
	return 
}

func main(){
	var src MyStringData
	src = MyStringData{
		str: "dhhsdsdhshdwdhwwd",
	}
	p := make([]byte, 4)
	fmt.Println("Len of src.str", len([]byte(src.str)))
	for {
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break;
		} else if err != nil {
			fmt.Println("something went wrong")
			break;
		}

	}
}

// continue to https://medium.com/rungo/introduction-to-streams-and-buffers-d148c0cda0ad