Hello Notes
===========

Notes regarding the compilation of simple "Hello, world!" apps.


| Language              | Category                               | Source Size | Binary Bytes | Binary Size | Compile Line                                                                        | Notes                                                 |
| --------              | --------                               | ----------: | -----------: | ----------: | ------------                                                                        | -----                                                 |
| Assembler (GAS)       | static string output                   | 458 B       | 8,328        | 8.1 KB      | `as hello_str_64.s -o hello_str_64.o && ld hello_str_64.o -e _main -o hello_str_64` |                                                       |
| Assembler (NASM)      | static string output                   | 1,199 B     | 8,288        | 8.1 KB      | `nasm -f macho hello_str.asm && ld -o hello_str -e mystart hello_str.o`             |                                                       |
| BASH                  | static string output                   | 41 B        | 41 + bash    | 41 B        | N/A                                                                                 | Requires the BASH shell                               |
| BASH                  | variable string output                 | 49 B        | 49 + bash    | 49 B        | N/A                                                                                 | Requires the BASH shell                               |
| BASH                  | argument string output                 | 48 B        | 48 + bash    | 48 B        | N/A                                                                                 | Requires the BASH shell                               |
| C                     | static string output                   | 73 B        | 8,432        | 8.2 KB      | `gcc -ansi hello_str.c -o hello_str`                                                |                                                       |
| C                     | static string output w/ncurses         | 118 B       | 8,676        | 8.5 KB      | `gcc -ansi hello_ncurses.c -o hello_ncurses -lcurses`                               |                                                       |
| C                     | variable string output                 | 94 B        | 8,432        | 8.2 KB      | `gcc -ansi hello_var.c -o hello_var`                                                |                                                       |
| C                     | argument string output                 | 102 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_arg.c -o hello_arg`                                                | `argv[1]` in the `printf`                             |
| C                     | argument variable string output        | 115 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_argvar.c -o hello_argvar`                                          | Pointer variable                                      |
| C                     | argument memcpy variable string output | 154 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_argvar_memcpy.c -o hello_argvar_memcpy`                            |                                                       |
| C                     | argument strcpy variable string output | 150 B       | 8,496        | 8.3 KB      | `gcc -ansi hello_argvar_strcpy.c -o hello_argvar_strcpy`                            |                                                       |
| C++ (CLang)           | static string output                   | 91 B        | 15,204       | 14.85 KB    | `clang++ -lstdc++ hello_str.cpp -o hello_str`                                       |                                                       |
| C++ (CLang Optimized) | static string output                   | 91 B        | 10,679       | 10.43 KB    | `clang++ -lstdc++ hello_str.cpp -O3 -o hello_str`                                   |                                                       |
| Go                    | minimalist static string output        | 55 B        | 1,070,112    | 1.1 MB      | `go build hello_min.go`                                                             | Didn't import the `fmt` package. Just used `println`. |
| Go                    | static string output                   | 73 B        | 2,253,776    | 2.3 MB      | `go build hello_str.go`                                                             |                                                       |
| Go                    | variable string output                 | 88 B        | 2,253,760    | 2.3 MB      | `go build hello_var.go`                                                             | `fmt.Printf` is smaller than `fmt.Println`?           |
| Go                    | arg string output                      | 94 B        | 2,253,760    | 2.3 MB      | `go build hello_arg.go`                                                             | `os.Args` package is free with `fmt` package?         |
| ooc                   | static string output                   | 25 B        | 612,244      | 597.90 KB   | `rock -v hello_str.occ`                                                             |                                                       |
| ooc                   | variable string output                 | 37 B        | 612,332      | 597.98 KB   | `rock -v hello_var.occ`                                                             |                                                       |
| ooc                   | arg string output                      | 64 B        | 616,484      | 602.04 KB   | `rock -v hello_arg.occ`                                                             |                                                       |
| Rust                  | static string output                   | 45 B        | 302,620      | 303 KB      | `rustc hello_str.rs`                                                                |                                                       |
| Rust                  | variable string output                 | 63 B        | 302,860      | 303 KB      | `rustc hello_var.rs`                                                                |                                                       |
| Rust                  | arg string output                      | 127 B       | 334,420      | 336 KB      | `rustc hello_arg.rs`                                                                |                                                       |
| Scheme                | static string output                   | 34 B        | 14,160       | 13.83 KB    | `csc hello_str.scm`                                                                 | Compiled using [CHICKEN Scheme][]                     |
| Scheme                | variable string output                 | 74 B        | 14,272       | 13.94 KB    | `csc hello_var.scm`                                                                 | Compiled using [CHICKEN Scheme][]                     |
| Scheme                | arg string output                      | 84 B        | 14,280       | 13.95 KB    | `csc hello_arg.scm`                                                                 | Compiled using [CHICKEN Scheme][]                     |



### Basic Finding

| Language  | Language Simplicity | Binary Size | Binary Speed | Notes |
| --------  | :-----------------: | :---------: | :----------: | ----- |
| Assembler | 5                   | 2           | 1            |       |
| BASH      | 1                   | 1           | 5            |       |
| C         | 4                   | 2           | 2            |       |
| Rust      | 3                   | 3           | 3            |       |
| Go        | 3                   | 5           | 4            |       |
| ooc       | 3                   | 4           |              |       |

* Assembler is the most verbose and thus complex language.
* Beyond this I find Go to be the simplest language to code in. But Rust is a close second.
* Assembler and C create the smaller binaries, by far.
* Go and Rust source generally simpler than C but VERY C based.
* Rust binaries are very large compared to Assemply/C binaries.
* Go binaries are enormous by any standard. But is the simplest language to code in.

Speed Runs
----------

I realize these are crap example of testing speed. The applications are too simple. But it still gives some reference for discussion.

Running `time ./hello_str` three times for each binary.

| Language                | real     | user     | sys      | Date       | Version                                                                      |
| --------                | ---:     | ---:     | --:      | :--:       | -------                                                                      |
| Assembler (GAS 64-bit)  | 0m0.002s | 0m0.001s | 0m0.001s | 2016-08-15 | Apple LLVM version 7.3.0 (clang-703.0.31)...                                 |
| Assembler (NASM 32-bit) | 0m0.001s | 0m0.000s | 0m0.001s | 2016-08-15 | NASM version 0.98.40 (Apple Computer, Inc. build 11) compiled on Feb 10 2016 |
| BASH                    | 0m0.022s | 0m0.002s | 0m0.002s | 2016-08-15 | GNU bash, version 4.3.46(1)-release (x86_64-apple-darwin15.5.0)...           |
| C                       | 0m0.009s | 0m0.001s | 0m0.002s |            |                                                                              |
| Go                      | 0m0.004s | 0m0.001s | 0m0.002s | 2016-08-15 | go version go1.7 darwin/amd64                                                |
| ooc                     | 0m0.004s | 0m0.002s | 0m0.002s | 2016-08-15 | rock 0.9.10 codename rita...                                                 |
| Rust                    | 0m0.003s | 0m0.001s | 0m0.001s | 2016-08-15 | rustc 1.10.0                                                                 |
| Scheme                  | 0m0.005s | 0m0.002s | 0m0.002s | 2016-08-15 | ...Version 4.11.0 (rev ce980c4)...                                           |




Running `time ./hello_arg RuneImp` three times for each binary.

| Language | real     | user     | sys      | Date       | Version                                                            |
| -------- | ---:     | ---:     | --:      | :--:       | -------                                                            |
| BASH     | 0m0.027s | 0m0.002s | 0m0.002s | 2016-08-15 | GNU bash, version 4.3.46(1)-release (x86_64-apple-darwin15.5.0)... |
| C        | 0m0.007s | 0m0.001s | 0m0.002s | original   |                                                                    |
| Go       | 0m0.019s | 0m0.001s | 0m0.004s | 2016-08-15 | go version go1.7 darwin/amd64                                      |
| ooc      | 0m0.004s | 0m0.002s | 0m0.002s | 2016-08-15 | rock 0.9.10 codename rita...                                       |
| Rust     | 0m0.003s | 0m0.001s | 0m0.001s | 2016-08-15 | rustc 1.10.0                                                       |
| Scheme   | 0m0.006s | 0m0.002s | 0m0.002s | 2016-08-15 | ...Version 4.11.0 (rev ce980c4)...                                 |





Running `time ./hello_str` three times for each binary.

| Language                | real     | user     | sys      | Date     | Version                       |
| --------                | ---:     | ---:     | --:      | :--:     | -------                       |
| Assembler (GAS 64-bit)  | 0m0.009s | 0m0.000s | 0m0.002s | original |                               |
| Assembler (NASM 32-bit) | 0m0.003s | 0m0.000s | 0m0.001s | original |                               |
| BASH                    | 0m0.037s | 0m0.002s | 0m0.006s | original | 4.?                           |
| C                       | 0m0.009s | 0m0.001s | 0m0.002s | original |                               |
| Go                      | 0m0.019s | 0m0.001s | 0m0.004s | original | go version go1.6 darwin/amd64 |
| ooc                     | 0m0.004s | 0m0.002s | 0m0.002s | original | rock 0.9.10 codename rita...  |
| Rust                    | 0m0.008s | 0m0.001s | 0m0.002s | original | 1.7.?                         |
| Scheme                  | 0m0.005s | 0m0.002s | 0m0.002s | original | ?                             |


Running `time ./hello_arg RuneImp` three times for each binary.

| Language | real     | user     | sys      | Date     | Version                       |
| -------- | ---:     | ---:     | --:      | :--:     | -------                       |
| BASH     | 0m0.022s | 0m0.002s | 0m0.004s | original | 4.?                           |
| C        | 0m0.007s | 0m0.001s | 0m0.002s | original |                               |
| Go       | 0m0.019s | 0m0.001s | 0m0.004s | original | go version go1.6 darwin/amd64 |
| ooc      | 0m0.004s | 0m0.002s | 0m0.002s | original | rock 0.9.10 codename rita...  |
| Rust     | 0m0.012s | 0m0.001s | 0m0.002s | original | 1.7.?                         |
| Scheme   | 0m0.006s | 0m0.002s | 0m0.002s | original | ?                             |


Compile Line
------------

Assembler (GAS)
: `as appname.s -o appname.o`
: `ld appname.o -e _main -o appname`
: Note that _main could be anything but is what's specified in the example source.

Assembler (NASM)
: `nasm -f macho appname.asm -o appname.o`
: `ld appname.o -e mystart -o appname`
: Note that mystart could be anything but is what's specified in the example source. _start is the default in NASM.

C
: `gcc -ansi appname.c -o appname`

Go
: `go build appname.go`

Rust
: `rustc appname.rs`


Tool Version
------------

BASH
: `bash --version`

C
: `gcc --version`

GAS
: `as --version`

Go
: `go version`

NASM
: `nasm -v`

ooc
: `rock -V`

Rust
: `rustc -V`

Scheme
: CHICKEN: `csc -version`

ToDo
----

* [ ] Setup proper testing
* [ ] Setup make file to build all of these
* [ ] Figure out how to code `hello_arg.asm`





[CHICKEN Scheme]: https://www.call-cc.org/