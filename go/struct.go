package main

import "fmt"

func TestStruct() {
	type myStruct struct {
		a bool
		b int64
		c string
	}

	test := myStruct{a: true}
	fmt.Print(test)

	testAnonim := struct {
		a bool
		b int64
		c string
	}{a: false}

	test = testAnonim
	fmt.Print(test)
}
