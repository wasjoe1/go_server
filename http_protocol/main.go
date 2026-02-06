package main

import (
	"fmt"
	"io"
	// "os" not required anymore since we stop reading from a file
	"log"
	"bytes"
	"net" // to set up TCP conn
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

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection has been accepted")
}

func main() {
	// task 4: refactor to have a function that reads lines from a TCP connection
	// filename := "messages.txt"
	// f, err := os.Open(filename)
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer ln.Close() // leaving it here as you can see where u "close" ur resource and know that this would be safe
	
	// previous code for handling lines from a file:
	// for line := range getLinesChannel(f) {
	// 	fmt.Printf("read: %s\n", line) // print the line to stdout
	// }

	for {
		conn, err := ln.Accept() // block until a connection is accepted
		if err != nil {
			log.Printf("Client connection error:", err) // dont use Fatal if not the currently opened socket will not be closed, causing resrc leak => program immediately terminates with os.Exit(1)
			continue // should not crash the server
		}

		for line := range getLinesChannel(conn) { // take in conn instead of f
			fmt.Printf("read: %s\n", line) // print the line to stdout
		}

		// go handleConnection(conn) // start an async goroutine that is immediately
	}

}
