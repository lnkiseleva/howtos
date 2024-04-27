package main

import "fmt"

func main() {
	num := []string{"1", "2", "3"}
	funcs := []func(){}
	for index, value := range num {
		fmt.Printf("[%d] - %s\n", index, value)
		funcs = append(funcs, func() { fmt.Print(value) })
	}

	for _, value := range funcs {
		value()
	}

	// TestStruct()

	// TestPointers()

	// TestSlice()

	// TestTaskWb()

	TestBrackets("dsd(sds)")
}
