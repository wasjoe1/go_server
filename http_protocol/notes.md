# Notes

## Main Notes
- why need init go module?
1. define the module's identity
2. manage dependencies for project
3. module context required to handle imports & integrate with GO toolchain features

- fmt.Fprintln vs fmt.Println
    fmt.Println("hello") => always prints to stdout `fd 0, 1 & 2 are for stdin, stdout & stderr by default opened by the OS`
    fmt.Fprintln(os.Stderr, "error:", err) => takes in any io.Writer, & the rest of the params are string args where a single space is inserted between them `func Fprintln(w io.Writer, a ...any) (n int, err error)`

- GO does not have F-string literals like python
    - From fastest → slowest:
        1. strings.Builder
        2. string + strconv
        3. fmt.Sprint
        4. fmt.Sprintf
    - string rules of thumb
        Debug / logs → fmt.Sprintf
        Fast paths → concatenation or strings.Builder
        Large strings → strings.Builder
        One-off prints → fmt.Println / fmt.Printf
- slice
    - not an array
    - is a view over an array => *UTH its implemented using a struct with the fields: data, length, capacity
        - data (or ptr) - a ptr to the first element of the underlying array that the slice can access
        - len - # of elements accessible in the slice
        - cap - max # of ele that underlying array can hold
    - i.e. `make([]T, len, cap)` => when only 1 size is speicified `make([]T, size)` len == cap
    ```go
    type slice struct {
        ptr *T
        len int
        cap int
    }
    ```
- printing to stdout (concatenation & format)
    format: * use fmt.printf when printing to stdout => format strings %s, %v etc.
	concatenation: * & fmt.Println => insert space betwn args, & \n char at concatenationend
*UTH - under the hood

- state machine
defn: A state machine is a mathematical abstraction used to design algorithms.
basically reads a set of inputs and changes to a different state, based on those inputs.
https://developer.mozilla.org/en-US/docs/Glossary/State_machine

- io.ReadCloser
it is an interface from the GO standard library
```go
type ReadCloser interface {
    Read(p []byte) (n int, err error)
    Close() error
}
```

- channels
`chan string` => send & receive channel, carries string type data
`<- chan string` => receive only channel, carries string type data; cant insert values to channel
    allowed action: value := <- ch // blocks until value is able to receive
`chan <- string` => send only channel, carries string type data; cant retrieve values from the channel
    allowed action: ch <- value // send value to ch; blocks until the other side is ready to receive

__unbuffered channels__:
ch := make(chan string)
Capacity = 0
ch <- value
* this means blocks until a receiver is ready, at most 1 value is stored (MEM usage constant)

__Buffered channel__
ch := make(chan string, 16)
Capacity = 16
Holds at most 16 strings => bounded memory with a small queue

__[]string ?__
why []string cant be used? wouldnt the exercise where we ingest data by chunks show that we can use string arrays initially and chunk them to not overrun the MEM?
No.
looking at this function `func getLines(f io.ReadCloser) []string`
this means the entire string needs to be read into MEM
whereas if we do `func getLines(f io.ReadCloser) <-chan string` the content can be buffered inside the channel first & block the socket from continuing to send us data

__how to add data into receive-only channel & vice versa__
chanel types are interpreted by the caller of the function, not inside the function it was created
```go 
func getLinesChannel(f io.ReadCloser) <-chan string {
    ch := make(chan string) // channel creation, currently its bi directional
        // you can send (ch <- value)
        // you can receive (val := <-ch)
        // you can close it 
    value := 3
    ch <- value
    return ch // GO implicitly converts `chan string` to `<-chan string`
}
```

- IIFE
```go
// this defines an annonymous function and immediately executes it asynchronously
go func() {
    ...
}()   // ← call
```
pattern is called an IIFE (Immediately Invoked Function Expression)

## Miscellaneous Notes/ Fun Facts
-  RFC (request for comments) - official technical doc that defines how the network protocol works
    - is a published doc by the IETF (internet engineering task force) => some org that standardises protocols that allow the internet to function LOL

### boot.dev cli installation
- this course required me to download the boot.dev CLI to run the test cases against my local env (refer to github to install @ https://github.com/bootdotdev/bootdev?tab=readme-ov-file#installation)
[GO-installation]
```bash
joechua@r-98-107-25-172 ~ % curl -sS https://webi.sh/golang | sh


>>> Welcome to Webi! - modern tools, instant installs.  <<<
    We expect your experience to be absolutely perfect!

    Success? Star it!   https://github.com/webinstall/webi-installers
    Problem? Report it: https://github.com/webinstall/webi-installers/issues
                        (your system is Darwin/arm64 with libc & curl)

Bootstrapping Webi
    Downloading https://webi.sh/packages/webi/webi.sh
        to ~/.local/bin/webi
    Running ~/.local/bin/webi golang@stable

Installing go ...
    Found  ~/.local/bin
    Initializing ~/.config/envman/
    Edit ~/.bashrc to source ~/.config/envman/load.sh
    Edit ~/.zshrc to source ~/.config/envman/load.sh
    WARN: possible PATH conflict between 'go1.25.6' and currently installed version
    /Users/joechua/.local/opt/go/bin/go (new)
    /usr/local/go/bin/go (existing)
    Downloading go from
      https://dl.google.com/go/go1.25.6.darwin-arm64.tar.gz
    Saved as ~/Downloads/webi/go/1.25.6/go1.25.6.darwin-arm64.tar.gz
    Extracting ~/Downloads/webi/go/1.25.6/go1.25.6.darwin-arm64.tar.gz
    Installing to ~/.local/opt/go-v1.25.6/bin/go
    Removing /Users/joechua/.local/opt/go-v1.25.6
    Moving go
      to ~/.local/opt/go-v1.25.6
Installed 'go v1.25.6' to ~/.local/opt/go

    Edit ~/.config/envman/PATH.env to add:
        ~/.local/bin
        ~/.local/opt/go/bin
        ~/go/bin

>>> ACTION REQUIRED <<<
        Copy, paste & run the following command:
        source ~/.config/envman/PATH.env
        (newly opened terminal windows will update automatically)

joechua@r-98-107-25-172 ~ % 
```
boot.dev installation
```bash
joechua@r-98-107-25-172 http_protocol % bootdev --version
bootdev version v1.23.3
joechua@r-98-107-25-172 http_protocol %
```
test submission
```bash
Please navigate to:
https://boot.dev/cli/login

Logged in successfully!
joechua@r-98-107-25-172 http_protocol % bootdev run b0cebf37-7151-48db-ad8a-0f9399f94c58
╭───────────────────────────────────────╮
│ Running: echo "bootdev cli is ready!" │
╰┬──────────────────────────────────────╯
 ├─  ?  Expect exit code 0
 ├─  ?  Expect stdout to contain all of:
 │       - 'bootdev cli is ready!'      

 > Command exit code: 0
 > Command stdout:

bootdev cli is ready!
joechua@r-98-107-25-172 http_protocol % bootdev run b0cebf37-7151-48db-ad8a-0f9399f94c58 -s
╭──────────────────────────────────────────╮
│ ✓  Running: echo "bootdev cli is ready!" │
╰┬─────────────────────────────────────────╯
 ├─  ✓  Expect exit code 0
 ├─  ✓  Expect stdout to contain all of:
 │       - 'bootdev cli is ready!'      

 > Command exit code: 0
 > Command stdout:

bootdev cli is ready!


All tests passed! 🎉

Return to your browser to continue with the next lesson.

joechua@r-98-107-25-172 http_protocol % 
```