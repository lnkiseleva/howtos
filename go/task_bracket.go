package main

import "fmt"

func TestBrackets(s string) {
	a := test(s)
	fmt.Println(a)
}

func test(s string) bool {
	count := 0
	for _, val := range s {
		str := string(val)
		if str == "(" {
			count++
		} else if str == ")" {
			count--
		}
		if count < 0 {
			return false
		}
	}

	return count == 0
}
