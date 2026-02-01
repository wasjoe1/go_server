package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "messages.txt"
	f, err := os.Open(filename)
	if (err != nil) {
		fmt.Fprintln(os.Stderr, "open:", err)
	}

	defer f.Close()

	buf := make([]byte, 8)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			fmt.Printf("read: %s\n", string(buf[:n])) // cleaner to convert to string imo
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error encountered:", err)
			return
		}
	}
}