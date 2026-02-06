package main

import (
	"io"
	"log"
	"bytes"
	"net" // to set up TCP conn
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	// read lines goroutine
	recv_chan := make(chan string, 1)
	read_all_lines := func() {
		defer close(recv_chan)
		buf := make([]byte, 8)
		str := ""
		for {
			n, err := f.Read(buf)
			
			if n > 0 {
				data := buf[:n]
				if i := bytes.IndexByte(data, '\n'); i != -1 {
					str += string(data[:i])
					recv_chan <- str
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
				log.Printf("error encountered:", err)
				return
			}
		}
		// print out last bit of string
		if len(str) != 0 {
			recv_chan <- str
		}
	}
	go read_all_lines()

	return recv_chan
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("Connection has been accepted")
}

func main() {

	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer ln.Close()
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Client connection error:", err)
			continue
		}

		for line := range getLinesChannel(conn) {
			log.Printf("read: %s\n", line)
		}
		// go handleConnection(conn) // temp code for future abstraction
	}

}
