package main

import (
	"fmt"
	"leandersonandre/dsa/pkg/linkedlist"
)

func main() {
	fmt.Println("Hello, World!")
	l := linkedlist.New()
	l.AddLast(0)
	l.AddFirst(1)
	l.AddFirst(2)
	l.AddFirst(3)
	l.AddLast(9)
	fmt.Println(l.Size())
	l.PrintList()
}