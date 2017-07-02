package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func applyRot13(c, offset int) int {
	return ((c-offset)+13)%26 + offset
}

func rot13(c byte) byte {
	n := c
	switch {
	case (c >= 65 && c <= 91):
		n = byte(applyRot13(int(c), 65))
	case (c >= 97 && c <= 122):
		n = byte(applyRot13(int(c), 97))
	}
	return n
}

// read the inner read stream and manipulate it
func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	if err == nil {
		for i := 0; i < n; i++ {
			b[i] = rot13(b[i])
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Println(s)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
