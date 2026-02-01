package main // defines an executable program; every GO src file starts with a package declaration => "package mathutils" defines a library

import "fmt" // stands for format; Go's std lib pckg for printing, scanning & string formatting

// any go executable program always needs a main function => entry point for an executable program in like in C
func main() {
	// Sad variable declaration: var style
	var smsSendingLimit int // declares the variable; GO then auto initializes it to the 0 value
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username) // 0 0.00 false ""
	// %v => works for any type
	// %f => float64
	// %.2f => exactly 2 digits after the floating decimal point
	// %q print quoted strings instead of just _ it will be "_"


	// Best way: walrus operator :=
		// declares a new variable & assigns a value to it in 1 line
	my_skill_issues := 42 // will be inferred as an int => type inference!
	// walrus is preferred over var style declarations; only bad thing is that walrus declarations cant be used outside of a funciton (in global/ package scope)

	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd) // pritns "Happy birthday! You are now 21 years old!"

	/*
	- package main lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
	- import "fmt" imports the fmt (formatting) package from the standard library. It allows us to use fmt.Println to print to the console.
	- func main() defines the main function, the entry point for a Go program.
	*/

	/*
	Two Kinds of Errors (in programming):
	1. Compilation errors. Occur when code is compiled. It's generally better to have compilation errors because they'll never accidentally make it into production. You can't ship a program with a compiler error because the resulting executable won't even be created.
	2. Runtime errors. Occur when a program is running. These are generally worse because they can cause your program to crash or behave unexpectedly.
	
	* i.e. compilation err could be wrong syntax causes code to not compile
	* i.e. run time could be doing 1/0 => program was running just that an err happened during runtime
	*/

	/*
	GO's speed 
	generally runs faster than interpreted languages, & compiles faster than other compiled languages => allows for more productive developer experience
	*/

	/*
	signed ints: int  int8  int16  int32  int64
	unsinged ints: uint uint8 uint16 uint32 uint64 uintptr
	signed decimals: float32 float64
	complex numbers: complex64 complex128 => has a real & imaginary part

	* the sizes above represent how many bits in mem will be used to store that variable
	* default int & uint types refer to either 32 or 64 bit sizes depending on the env of the user
	* standard sizes are:
		- int
		- uint
		- float64
		- complex128
	* other default types
		- bool
		- string
		- byte
		- rune
	* straying from default types can make code messy => only use specific types for performance (time & MEM)
	i.e. bad code
	var myAge uint16 = 25
	myAgeInt := int(myAge) // returns the int type => platform dependent so could be 32 or 64
	*/

	// converting types
	temperatureFloat := 88.26
	temperatureInt := int64(temperatureFloat) // casting a float to int => truncates the floating point portion; becomes 88

	mileage, company := 80276, "Toyota" // same as the below 2 lines
	mileage := 80276
	company := "Toyota"

	/*
	GO Run time - run time library that is part of every GO program
	it implements garbage collection, concurrency, stack management & other critical features etc.
	is anagolous to libc (C stadard library)
	* is a small amount of extra code included in the executable binary
	*/

	// constants
	const pi = 3.14159 // can only be primitive types i.e. strings, integers, booleans & floats
	// cant be slices, maps & structs
	// *IMPT: constants must be known at compile time => computation must happen during compile time
	// valid computation
	const firstName = "Lane"
	const lastName = "Wagner"
	const fullName = firstName + " " + lastName // i think, think of it as the compiler doing the actual computation

	// invalid computaitno
	// the current time can only be known when the program is running
	const currentTime = time.Now() // fails computation; valid in JS but not GO

	/*
	Go is generally faster and more lightweight than interpreted or VM-powered languages like:
	- Python
	- JavaScript
	- PHP
	- Ruby
	- Java

	still lgas behind:
	- C
	- C++
	- Rust
	
	* GO is slower due to its automated MEM mgmt aka GO runtime => slower speed but mem safe & simple syntax
	*/

	// *IMPT!! ok lowkey i stopped here since i see that they are just going through all the basic programming language stuff
	// *IMPT!! i did see 1 thing unique about the GO language is channels and that is a specific concept i should pick up

	// writing a function in GO
	
	// params share the same type
	func multiply(x int, y int) int { // return type is after the parantheses
 	   return x * y
	}
	func multiply(x, y int) int {
 	   return x * y
	}
	// no return value
	func logMessage(msg string) {
		fmt.Println(msg)
	}
	// multiple return values
	func divide(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	}
	// using multiple values in GO
	result, _ := divide(10, 2) // if you want to ignore errors
	result, err := divide(10,2)
	result := divide(10,2) // !!THROWS a compile time error!!
	// * take note that this is not a tuple and that its specific that this function then requires 2 variables to store the returned values

	// GO goroutines
	func handleRequests(reqs <-chan request) {
		// changing it to use go routines
		for req := range reqs {
			go handleRequest(req) // go keyword runs the function in a lightweight thread
		}
	}
	// 1. goroutines is a concurrent unit of execution, allowing functions to run concurrently
	// 2. all programs start with main goroutine
	// 3. NOT OS or green threads; coroutines that integrate with GO runtime; rely on GO runtime to suspend them/
	// 4. non preemptive (OS cant force them to context switch, unlike threads)
	// 5. (a) do all goroutines run on 1 thread? NO
	// 	  (b) does each goroutine run on 1 thread? NO
	// 		=> Go uses an (M:N) scheduler, where M # of goroutines multiplex on a much smaller num of N OS threads
}