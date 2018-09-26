package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image ...
type Image struct {
	w int
	h int
}

// ColorModel ...
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds ...
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At ...
func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y) // or v := uint8(x - y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
