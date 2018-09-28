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
