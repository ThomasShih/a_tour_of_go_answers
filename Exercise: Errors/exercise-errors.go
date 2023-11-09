package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 1 {
		return 0, ErrNegativeSqrt(x)
	}

	var z float64 = 1 // Recommended starting point

	// Newton's method
	var last_z float64 = z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		// Stop if z is not changing
		if math.Abs(last_z-z) < 0.000001 {
			break
		} else {
			last_z = z
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
