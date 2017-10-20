package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	if err == nil {
		for i := range b {
			r := b[i]
			if strings.ContainsRune("ABCDEFGHIGKLMabcdefghijklm", rune(r)) {
				b[i] = b[i] + 13
			}
			if strings.ContainsRune("NOPQRSTUVWXYZnopqrstuvwxyz", rune(r)) {
				b[i] = b[i] - 13
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
