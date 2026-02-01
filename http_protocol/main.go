package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"bytes"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	// TODO: takes in an arg that is readable & closeable

	// read from f
	// chunk data myself? chunk into lines
	// ch <- line

	recv_chan := make(chan string, 1) // create the channel string type with buffer 1 string
	// func read_all_lines() {
	read_all_lines := func() {
		defer close(recv_chan) // defer  closing the channel until all lines have been fully read & receiver has received
		buf := make([]byte, 8)
		str := ""
		for {
			n, err := f.Read(buf)
			
			if n > 0 {
				data := buf[:n]
				if i := bytes.IndexByte(data, '\n'); i != -1 {
					str += string(data[:i])
					recv_chan <- str // fmt.Printf("read: %s\n", str)
					str = string(data[i+1:])
				} else {
					str += string(data)
				}
			}
			// EOF
			if err == io.EOF {
				break
			}
			// Random error
			if err != nil {
				fmt.Printf("error encountered:", err)
				return
			}
		}
		// print out last bit of string that doesnt end with \n
		if len(str) != 0 {
			recv_chan <- str // fmt.Printf("read: %s\n", str)
		}
	}

	go read_all_lines()
	// this will get blocked once the buffer in the channel is full,
	// allowing the function to continue & return recv_chan allowing the user to range over it
	// if we dont asynchronously do it, it will just get blocked & stucked and getLinesChannel will never return

	return recv_chan
}


func main() {
	// task 4: refactor to have a function that reads lines from a TCP connection
	filename := "messages.txt"
	f, err := os.Open(filename)
	if (err != nil) {
		log.Fatal()
	}
	defer f.Close() // leaving it here as you can see where u "close" ur resource and know that this would be safe
	
	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line) // print the line to stdout
	}
}
