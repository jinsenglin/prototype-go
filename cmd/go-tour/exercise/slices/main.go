package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, 0, dy)
	for i := 0; i < dy; i++ {
		x := make([]uint8, 0, dx)
		for j := 0; j < dx; j++ {
			x = append(x, uint8(0))
		}
		y = append(y, x)
	}
	return y
}

func main() {
	pic.Show(Pic)
}
