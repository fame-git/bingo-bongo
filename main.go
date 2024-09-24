package main

import (
	"fmt"
	"io"
	"log"
)

type SlowReader struct {
	contents string
	pos      int
}

func (m *SlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.contents) {
		n := copy(p, m.contents[m.pos:m.pos+1])
		// 0:1 => h
		// 1:2 => e

		// 0 <= 1
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {

	slowReaderInstance := &SlowReader{
		contents: "hello world!",
	}

	out, err := io.ReadAll(slowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s\n", out)
}
