package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(f ErrNegativeSqrt) (float64, error) {
	if f < 0 {
		return 0, f
	}
	x := float64(f)
	z1 := float64(1)
	z2 := float64(100)
	for z1 != z2 {
		var tmp = z2
		z2 = z1 - (z1*z1-x)/(2*z1)
		z1 = tmp
	}
	return z2, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
