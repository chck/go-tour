package main

import "fmt"

// fibonacci is a function that returns
// a functino that returns an int.
func fibonacci() func() int {
	n0, n1 := 1, 1
	return func() (res int) {
		res = n0
		n0 = n1
		n1 = n0 + n1
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
