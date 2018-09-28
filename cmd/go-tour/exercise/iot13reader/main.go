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
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	// NOTE: here can't use `for range` !!!

	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 'A' && b[i] <= 'M':
			b[i] = b[i] + 13
		case b[i] >= 'N' && b[i] <= 'Z':
			b[i] = b[i] - 13
		case b[i] >= 'a' && b[i] <= 'm':
			b[i] = b[i] + 13
		case b[i] >= 'n' && b[i] <= 'z':
			b[i] = b[i] - 13
		default:
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
