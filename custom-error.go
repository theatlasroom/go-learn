package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, err
	}
	closeEnough := true
	z := 1.0
	i := 0
	for closeEnough {
		zn := (z - ((z*z)-x)/(2*z))
		//fmt.Println(zn, z, (zn - z))
		delta := zn - z

		if delta < 0 {
			delta *= -1
		}

		if delta < 0.01 {
			closeEnough = false
		}

		z = zn
		i++
	}
	return z, nil
}

// define our ErrNegativeSqrt error
type ErrNegativeSqrt float64

// implement the error interface
func (e ErrNegativeSqrt) Error() string {
	val := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %s", val)
}

func main() {
	fmt.Println(Sqrt(-2))
}
