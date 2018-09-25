package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	last_squre := 0.0

	for i := 1; i < 11; i++ {
		z_squre := z * z

		fmt.Printf("#%d z=%g z*z=%g\n", i, z, z_squre)

		switch {
		case z_squre == x:
			return z, nil
		case z_squre > last_squre && z_squre-last_squre < 0.000000000000001:
			return z, nil
		case last_squre > z_squre && last_squre-z_squre < 0.000000000000001:
			return z, nil
		default:
			last_squre = z_squre
			z -= (z*z - x) / (2 * z)
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
