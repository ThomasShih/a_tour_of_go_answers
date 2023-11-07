package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
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
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
