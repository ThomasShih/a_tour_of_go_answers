package main

import "golang.org/x/tour/pic"

func gradient(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// Assign colors to each pixel.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			a[i][j] = gradient(i, j)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
