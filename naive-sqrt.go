package hello

import (
	"fmt"
)

func Sqrt(x float64) (float64, int) {
	closeEnough := true
	z := 1.0
	i := 0
	for closeEnough {
		zn := (z - ((z*z)-x)/(2*z))
		fmt.Println(zn, z, (zn - z))
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
	return z, i
}

func main() {
	result, iterations := Sqrt(17056595585)
	fmt.Println(result, iterations)
}
