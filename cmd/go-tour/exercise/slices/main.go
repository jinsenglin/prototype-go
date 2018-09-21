package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy, dy)
	for i := 0; i < dy; i++ {
		x := make([]uint8, dx, dx)
		y[i] = x
	}
	return y
}

func main() {
	pic.Show(Pic)
}
