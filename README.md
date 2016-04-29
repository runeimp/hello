Hello Notes
===========

Notes regarding the compilation of simple "Hello, world!" apps.


| Language         | Category                               | Source Size | Binary Bytes | Binary Size | Compile Line                                                                        | Notes                                                 |
| --------         | --------                               | ----------: | -----------: | ----------: | ------------                                                                        | -----                                                 |
| Assembler (GAS)  | static string output                   | 458 B       | 8,328        | 8.1 KB      | `as hello_str_64.s -o hello_str_64.o && ld hello_str_64.o -e _main -o hello_str_64` |                                                       |
| Assembler (NASM) | static string output                   | 1,199 B     | 8,288        | 8.1 KB      | `nasm -f macho hello_str.asm && ld -o hello_str -e mystart hello_str.o`             |                                                       |
| BASH             | static string output                   | 41 B        | 41+          | 41 B        | N/A                                                                                 | Requires the BASH shell                               |
| BASH             | variable string output                 | 49 B        | 49+          | 49 B        | N/A                                                                                 | Requires the BASH shell                               |
| BASH             | argument string output                 | 48 B        | 48+          | 48 B        | N/A                                                                                 | Requires the BASH shell                               |
| C                | static string output                   | 73 B        | 8,432        | 8.2 KB      | `gcc -ansi hello_str.c -o hello_str`                                                |                                                       |
| C                | variable string output                 | 94 B        | 8,432        | 8.2 KB      | `gcc -ansi hello_var.c -o hello_var`                                                |                                                       |
| C                | argument string output                 | 102 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_arg.c -o hello_arg`                                                | `argv[1]` in the `printf`                             |
| C                | argument variable string output        | 115 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_argvar.c -o hello_argvar`                                          | Pointer variable                                      |
| C                | argument memcpy variable string output | 154 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_argvar_memcpy.c -o hello_argvar_memcpy`                            |                                                       |
| C                | argument strcpy variable string output | 150 B       | 8,496        | 8.3 KB      | `gcc -ansi hello_argvar_strcpy.c -o hello_argvar_strcpy`                            |                                                       |
| Rust             | static string output                   | 45 B        | 302,620      | 303 KB      | `rustc hello_var.rs`                                                                |                                                       |
| Rust             | variable string output                 | 63 B        | 302,860      | 303 KB      | `rustc hello_var.rs`                                                                |                                                       |
| Rust             | arg string output                      | 127 B       | 334,420      | 336 KB      | `rustc hello_arg.rs`                                                                |                                                       |
| Go               | minimalist static string output        | 55 B        | 1,070,112    | 1.1 MB      | `go build hello_min.go`                                                             | Didn't import the `fmt` package. Just used `println`. |
| Go               | static string output                   | 73 B        | 2,253,776    | 2.3 MB      | `go build hello_str.go`                                                             |                                                       |
| Go               | variable string output                 | 88 B        | 2,253,760    | 2.3 MB      | `go build hello_var.go`                                                             | `fmt.Printf` is smaller than `fmt.Println`?           |
| Go               | arg string output                      | 94 B        | 2,253,760    | 2.3 MB      | `go build hello_arg.go`                                                             | `os.Args` package is free with `fmt` package?         |


### Basic Finding

| Language  | Language Simplicity | Binary Size | Binary Speed | Notes |
| --------  | :-----------------: | :---------: | :----------: | ----- |
| Assembler | 5                   | 2           | 1            |       |
| BASH      | 1                   | 1           | 5            |       |
| C         | 4                   | 2           | 2            |       |
| Rust      | 3                   | 3           | 3            |       |
| Go        | 3                   | 5           | 4            |       |

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

| Language                | real     | user     | sys      |
| --------                | ---:     | ---:     | --:      |
| Assembler (GAS 64-bit)  | 0m0.009s | 0m0.000s | 0m0.002s |
| Assembler (NASM 32-bit) | 0m0.003s | 0m0.000s | 0m0.001s |
| BASH                    | 0m0.037s | 0m0.002s | 0m0.006s |
| C                       | 0m0.009s | 0m0.001s | 0m0.002s |
| Go                      | 0m0.019s | 0m0.001s | 0m0.004s |
| Rust                    | 0m0.008s | 0m0.001s | 0m0.002s |

Running `time ./hello_arg RuneImp` three times for each binary.

| Language                | real     | user     | sys      |
| --------                | ---:     | ---:     | --:      |
| C                       | 0m0.007s | 0m0.001s | 0m0.002s |
| Rust                    | 0m0.012s | 0m0.001s | 0m0.002s |
| Go                      | 0m0.019s | 0m0.001s | 0m0.004s |
| BASH                    | 0m0.022s | 0m0.002s | 0m0.004s |


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


ToDo
----

* [ ] Setup proper testing
* [ ] Setup make file to build all of these
* [ ] Figure out how to code `hello_arg.asm`

