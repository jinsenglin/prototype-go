//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

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
