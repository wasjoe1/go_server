//go:build ignore
// +build ignore
/*
	OLD_MAIN.GO

	This file contains the code from previous tasks to store my learning
	i dont want to compile it => GO compiles and all `.go` files in the same dir & packg by default
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// task 1. print statement
	// fmt.Println("I hope I get the job!")

	// task2. read & write data streams
	// read text from messages.txt 8bytes per chunk, then print(stream) data to stdout 8 bytes per chunk as well

	// open file
	// while file still has text => no EOF
		// read 8 bytes of data
		// print 8 bytes of data
	filename := "messages.txt"
	f, err := os.Open(filename)
	// if there is an error
	if (err != nil) {
		fmt.Fprintln(os.Stderr, "open:", err)
	}
	defer f.Close() //to run this funciton when the surrounding function returns => clean up fd

	// example for loop in GO
	// for i:= 0; i < n; i++ {...}
	// i++ is a statement not expression! => does not return anything
	buf := make([]byte, 8) // 8 bytes chunk
	for { // GO only has 1 loop keyword "for"; for condition {} => dont put paranthese over the condition!!
		// while true, keep looping until EOF is encountered
		n, err := f.Read(buf) // read context into the bytes buffer, n is the length of the content
		// if we re-use the buffer, only the content up till length n is overwritten, everything else remains the same
		// buf = [ w x y z e f g h ]
		// 		↑ ↑ ↑ ↑
		// 	overwritten
		if n > 0 {
			// option 1: use formatted file print
			fmt.Fprintf(os.Stdout, "read: %s\n", buf[:n]) // %s accepts either byte slice ([]byte) OR string types => GO treats []byte as raw text data often
			// also, GO is quite strict on capitalisation => fmt.fprintf will throw an error

			//option 2: explicitly build string
			// toPrint := "read: " + string(buf[:n]) + "\n"
			// os.Stdout.Write([]byte(toPrint))

			// option 1 is cleaner and no explicit conversion		
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "read:", err)
			return // rmb to return/ break since it keeps looping!!
		}
	}
	
	// --------------------------------------------------------------------------------------------
	// upon review, sample code from boot.dev:
	for {
		b := make([]byte, 8, 8) // it might be safer to keep creating new buffers to ensure no old data is inside it. however its also the consideration of efficiency
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(b[:n]) // actually yea, this seems like a good practice
		fmt.Printf("read: %s\n", str)
	}
	// * use fmt.printf when printing to stdout => format strings %s, %v etc.
	// * & fmt.Println => insert space betwn args, & \n char at the end
}