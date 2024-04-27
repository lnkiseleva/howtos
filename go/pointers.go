package main

import "fmt"

type user struct {
	name string
}

func TestPointers() {
	a1 := 1
	b1 := &a1
	fmt.Print(a1, b1)

	a()
	b()
	c()

	fmt.Println(&a1)
	increment(&a1)
	u := escapeToHeap()
	fmt.Println(*u)
}

func increment(n *int) {
	*n++
	fmt.Println(*n, &n, n)
}

func escapeToHeap() *user {
	u := user{name: "test"}
	return &u
}

func a() {
	fmt.Print("1")
}

func b() {
	fmt.Print("2")
}

func c() {
	fmt.Print("3")
}
