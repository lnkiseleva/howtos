package main

import "fmt"

func TestSlice() {
	a := make([]int, 5, 10)

	for i, val := range a {
		fmt.Println(i, val)
	}
}
