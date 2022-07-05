package main

import (
	"fmt"
)

func main() {
	// a := [...]int{20,30,40,50,60,70,80,90}
	// s := a[2:5]
	// printSlice(s)
	// s2 := a[1:]
	// printSlice(s2)
	// copy(s,s2)
	// printSlice(s)

	
	slice:= make([]string, 2, 3) 
	func(slice []string){
			slice = append(slice, "a") 
			printSlice(slice)
			slice[0]="b";
			slice[1]="b"; 
			printSlice(slice)
	}(slice)
	printSlice(slice)


}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d address=%p %v\n", len(s), cap(s), &s[0], s)
}