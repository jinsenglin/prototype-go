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
