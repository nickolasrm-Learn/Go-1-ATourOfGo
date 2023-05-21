package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	for i := 0; i < n; i++ {
		v := p[i]
		switch {
		case v < 'A':
			continue
		case v > 'z':
			continue
		case v >= 'A' && v <= 'M' || v >= 'a' && v <= 'm':
			v += 13
		default:
			v -= 13
		}
		p[i] = v
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
