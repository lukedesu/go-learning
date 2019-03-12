package main

import "fmt"

// normal way
var age int

// batch way
var (
	name    string
	address string
	scores  []int
	d       func() bool
	e       struct {
		x int
	}
)

func main() {
	fmt.Println(e)
}
