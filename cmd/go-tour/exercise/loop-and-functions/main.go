package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	last_squre := 0.0

	for i := 1; i < 11; i++ {
		z_squre := z * z

		fmt.Printf("#%d z=%g z*z=%g\n", i, z, z_squre)

		switch {
		case z_squre == x:
			return z
		case z_squre > last_squre && z_squre-last_squre < 0.000000000000001:
			return z
		case last_squre > z_squre && last_squre-z_squre < 0.000000000000001:
			return z
		default:
			last_squre = z_squre
			z -= (z*z - x) / (2 * z)
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
