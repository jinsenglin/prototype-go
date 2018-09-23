package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1 := 0
	n2 := 1
	return func() int {
		n3 := n1 + n2
		n := n1
		n1 = n2
		n2 = n3
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
