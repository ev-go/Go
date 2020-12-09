package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	g := strings.NewReader("1234 5678 9abcd")

	b := make([]byte, 9)
	for {
		n, err := r.Read(b)

		fmt.Printf("n = %v; err = %v; b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}

		fmt.Printf("Next!\n")

		p, err := g.Read(b)
		fmt.Printf("n = %v; err = %v; b = %v\n", p, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
