# Notes

- GO is a compiled language (like C)
    - GO src code is compile to -> native machine code
    - no interpereter at run time (like python)
    - no VM (like JAVA's JVM which runs java byte code)
    - output is an executable file
- running GO programs => with your terminal & GO file
    run `go run main.go` [auto]
    1. GO compiles your code
    2. creates a temp. executable
    3. runs the executable
    4. deletes the executable
    run `go build main.go` [compile-only]
    1. after running, produces an executable ./main

- go features:
    - compiled => produces an executable
    - static type => checking of types at compile time
        - does seem strongly type (refer to below)
    - easily concurrent
    - has garbage collection
    - fast & lightweight

* when doing int * float this err showed up:
./main.go:8:15: invalid operation: costPerMessage * numMessagesFromDoris (mismatched types float64 and int) => seems to me this is a strongly typed language