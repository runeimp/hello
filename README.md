Hello Notes
===========

Notes regarding the compilation of simple "Hello, world!" apps.

| Language                      | Category                           | Source Size | Binary Bytes | Binary Size | Notes                                                 |
| --------                      | --------                           | ----------: | -----------: | ----------: | -----                                                 |
| Assembler (GAS - LLVM v8.0.0) | static string output               | 458 B       | 8,328        | 8.1 KB      |                                                       |
| Assembler (NASM v0.98.40)     | static string output               | 1,199 B     | 8,288        | 8.1 KB      |                                                       |
| BASH v4.4.5                   | static string output               | 41 B        | 41 + bash    | 41 B        | Requires the BASH shell                               |
| BASH v4.4.5                   | variable string output             | 49 B        | 49 + bash    | 49 B        | Requires the BASH shell                               |
| BASH v4.4.5                   | argument string output             | 48 B        | 48 + bash    | 48 B        | Requires the BASH shell                               |
| BASH v4.4.5                   | environment variable string output | 52 B        | 48 + bash    | 52 B        |                                                       |
| C (LLVM v8.0.0)               | static string output               | 73 B        | 8,432        | 8.23 KB     |                                                       |
| C (LLVM v8.0.0)               | static string output w/ncurses     | 118 B       | 8,668        | 8.46 KB     |                                                       |
| C (LLVM v8.0.0)               | variable string output             | 94 B        | 8,564        | 8.36 KB     |                                                       |
| C (LLVM v8.0.0)               | argument string output             | 115 B       | 8,432        | 8.23 KB     | `argv[1]` in the `printf`                             |
| C (LLVM v8.0.0)               | argument memcpy string output      | 154 B       | 8,564        | 8.36 KB     |                                                       |
| C (LLVM v8.0.0)               | argument strcpy string output      | 150 B       | 8,620        | 8.42 KB     |                                                       |
| C (LLVM v8.0.0)               | environment variable string output | 175 B       | 8,480        | 8.28 KB     |                                                       |
| C++ (CLang - LLVM v8.0.0)     | static string output               | 91 B        | 15,204       | 14.85 KB    |                                                       |
| C++ (CLang - LLVM v8.0.0)     | static string output               | 91 B        | 10,676       | 10.43 KB    |                                                       |
| Crystal v0.20.4               | static string output               | 20 B        | 105,968      | 103.48 KB   |                                                       |
| Crystal v0.20.4               | variable string output             | 31 B        | 105,968      | 103.48 KB   |                                                       |
| Crystal v0.20.4               | arg variable string output         | 36 B        | 106,040      | 103.56 KB   |                                                       |
| Crystal v0.20.4               | environment variable string output | 41 B        | 106,064      | 103.58 KB   |                                                       |
| D v2.072.2                    | static string output               | 65 B        | 933,008      | 911.14 KB   |                                                       |
| D v2.072.2                    | variable string output             | 95 B        | 1,017,016    | 993.18 KB   |                                                       |
| D v2.072.2                    | arg variable string output         | 97 B        | 1,016,992    | 993.16 KB   |                                                       |
| D v2.072.2                    | environment variable string output | 133 B       | 2,077,636    | 1.98 MB     |                                                       |
| Fantom v1.0.69                | static string output               | 84 B        | 84 + fan     | 84 B        |                                                       |
| Fantom v1.0.69                | variable string output             | 98 B        | 98 + fan     | 98 B        |                                                       |
| Fantom v1.0.69                | arg variable string output         | 129 B       | 129 + fan    | 129 B       |                                                       |
| Fantom v1.0.69                | environment variable string output | 133 B       | 133 + fan    | 133 B       |                                                       |
| Go v1.7.4                     | minimalist static string output    | 55 B        | 971,584      | 948.81 KB   | Didn't import the `fmt` package. Just used `println`. |
| Go v1.7.4                     | static string output               | 73 B        | 1,611,712    | 1.54 MB     |                                                       |
| Go v1.7.4                     | variable string output             | 89 B        | 1,620,096    | 1.55 MB     | `fmt.Printf` is smaller than `fmt.Println`?           |
| Go v1.7.4                     | arg variable string output         | 122 B       | 1,620,128    | 1.55 MB     | `os.Args` package is free with `fmt` package?         |
| Go v1.7.4                     | environment variable string output | 130 B       | 1,624,256    | 1.55 MB     |                                                       |
| ooc (rock 0.9.10)             | static string output               | 25 B        | 612,580      | 598.22 KB   |                                                       |
| ooc (rock 0.9.10)             | variable string output             | 37 B        | 612,668      | 598.31 KB   |                                                       |
| ooc (rock 0.9.10)             | arg variable string output         | 72 B        | 616,820      | 602.36 KB   |                                                       |
| ooc (rock 0.9.10)             | environment variable string output | 62 B        | 616,924      | 602.47 KB   |                                                       |
| Python 2.7.13                 | static string output               | 44 B        | 44 + python  | 44 B        |                                                       |
| Python 2.7.13                 | variable string output             | 58 B        | 58 + python  | 58 B        |                                                       |
| Python 2.7.13                 | arg variable string output         | 74 B        | 74 + python  | 74 B        |                                                       |
| Python 2.7.13                 | environment variable string output | 81 B        | 81 + python  | 81 B        |                                                       |
| Python 3.6.0                  | static string output               | 46 B        | 46 + python3 | 46 B        |                                                       |
| Python 3.6.0                  | variable string output             | 66 B        | 66 + python3 | 66 B        |                                                       |
| Python 3.6.0                  | arg variable string output         | 83 B        | 83 + python3 | 83 B        |                                                       |
| Python 3.6.0                  | environment variable string output | 89 B        | 89 + python3 | 89 B        |                                                       |
| Rust (rustc 1.14.0)           | static string output               | 44 B        | 346,504      | 338.38 KB   |                                                       |
| Rust (rustc 1.14.0)           | variable string output             | 62 B        | 346,696      | 338.57 KB   |                                                       |
| Rust (rustc 1.14.0)           | arg variable string output         | 123 B       | 379,784      | 370.88 KB   |                                                       |
| Rust (rustc 1.14.0)           | environment variable string output | 134 B       | 355,088      | 346.77 KB   |                                                       |
| Scheme (Chicken 4.11.0)       | static string output               | 34 B        | 14,160       | 13.83 KB    | Compiled using [CHICKEN Scheme][]                     |
| Scheme (Chicken 4.11.0)       | variable string output             | 74 B        | 14,272       | 13.94 KB    | Compiled using [CHICKEN Scheme][]                     |
| Scheme (Chicken 4.11.0)       | arg variable string output         | 84 B        | 14,280       | 13.95 KB    | Compiled using [CHICKEN Scheme][]                     |
| Scheme (Chicken 4.11.0)       | environment variable string output | 88 B        | 14,232       | 13.90 KB    | Compiled using [CHICKEN Scheme][]                     |
| Tcl 8.5.9                     | static string output               | 42 B        | 42 + tclsh   | 42 B        |                                                       |
| Tcl 8.5.9                     | variable string output             | 52 B        | 52 + tclsh   | 52 B        |                                                       |
| Tcl 8.5.9                     | arg variable string output         | 53 B        | 53 + tclsh   | 53 B        |                                                       |
| Tcl 8.5.9                     | environment variable string output | 48 B        | 48 + tclsh   | 48 B        |                                                       |


| Str Bytes | Language         |
| --------: | --------         |
| `0000041` | BASH             |
| `0000042` | Tcl              |
| `0000044` | Python 2         |
| `0000058` | Python 3         |
| `0000084` | Fantom           |
| `0008288` | Assembler (NASM) |
| `0008328` | Assembler (GAS)  |
| `0008432` | C                |
| `0010679` | C++              |
| `0014160` | Scheme           |
| `0105968` | Crystal          |
| `0346504` | Rust             |
| `0612580` | ooc              |
| `0933008` | D                |
| `1611712` | Go               |



### Basic Finding

| Language  | Language Simplicity | Binary Size | Binary Speed | Total | Simplicity & Speed |
| --------  | :-----------------: | :---------: | :----------: | :---: | :----------------: |
| Assembler | 5                   | 2           | 1            | 8     | 6                  |
| BASH      | 1                   | 1           | 1            | 3     | 6                  |
| C         | 4                   | 2           | 1            | 7     | 5                  |
| C++       | 3                   | 3           | 1            | 7     | 4                  |
| Crystal   | 2                   | 4           | 1            | 7     | 3                  |
| D         | 2                   | 5           | 1            | 8     | 3                  |
| Fantom    | 2                   | 1           | 4            | 7     | 6                  |
| Go        | 2                   | 5           | 1            | 8     | 3                  |
| ooc       | 1                   | 4           | 2            | 7     | 3                  |
| Python 2  | 1                   | 1           | 2            | 4     | 3                  |
| Python 3  | 1                   | 1           | 2            | 4     | 3                  |
| Rust      | 3                   | 4           | 1            | 8     | 4                  |
| Scheme    | 2                   | 3           | 1            | 6     | 3                  |
| Tcl       | 1                   | 1           | 2            | 3     | 3                  |



| Rank | Language Simplicity               | Binary Size | Binary Speed    |
| :--: | :-----------------:               | :---------: | :----------:    |
| 1    | Very simple for most basic tasks  | < 1 KB      | < 0m0.003s real |
| 2    | Simple for most basic tasks       | < 10 KB     | < 0m0.005s real |
| 3    | OK for most basic tasks           | < 100 KB    | < 0m0.007s real |
| 4    | Complex for most basic tasks      | < 1 MB      | < 0m0.009s real |
| 5    | Very complex for most basic tasks | >= 1 MB     | > 0m0.008s real |


The following thoughts were regarding Assembler, C, Go, and Rust only when I wrote them.

* Assembler is the most verbose and thus complex language.
* Beyond this I find Go to be the simplest language to code in. But Rust is a close second.
* Assembler and C create the smaller binaries, by far.
* Go and Rust source generally simpler than C but VERY C based.
* Rust binaries are very large compared to Assemply/C binaries.
* Go binaries are enormous by any standard. But is the simplest language to code in.


Speed Runs
----------

I realize these are crap example of testing speed. The applications are too simple. But it still gives some reference for discussion.

Running `time ./hello_str` five	times for each binary and averaging the results.

| Language                | real     | user     | sys      | Date       | Version                                                                      |
| --------                | ---:     | ---:     | --:      | :--:       | -------                                                                      |
| Assembler (GAS 64-bit)  | 0m0.003s | 0m0.001s | 0m0.001s | 2017-01-21 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                               |
| Assembler (NASM 32-bit) | 0m0.002s | 0m0.000s | 0m0.001s | 2017-01-21 | NASM version 0.98.40 (Apple Computer, Inc. build 11) compiled on Nov 15 2016 |
| BASH                    | 0m0.005s | 0m0.002s | 0m0.002s | 2017-01-21 | GNU bash, version 4.4.5(1)-release (x86_64-apple-darwin15.6.0)...            |
| C                       | 0m0.002s | 0m0.001s | 0m0.001s | 2017-01-21 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                               |
| C++                     | 0m0.002s | 0m0.001s | 0m0.001s | 2017-01-21 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                               |
| Crystal                 | 0m0.004s | 0m0.002s | 0m0.002s | 2017-01-21 | Crystal 0.20.4 (2017-01-06)                                                  |
| D                       | 0m0.003s | 0m0.001s | 0m0.001s | 2017-01-21 | DMD64 D Compiler v2.072.2                                                    |
| Fantom                  | 0m1.015s | 0m2.013s | 0m0.154s | 2017-01-21 | ...fan.version: 1.0.69...                                                    |
| Go                      | 0m0.003s | 0m0.000s | 0m0.002s | 2017-01-21 | go version go1.7.4 darwin/amd64                                              |
| ooc                     | 0m0.004s | 0m0.002s | 0m0.002s | 2017-01-21 | rock 0.9.10 codename rita, built on Sun Apr 10 01:50:13 2016                 |
| Python 2                | 0m0.027s | 0m0.010s | 0m0.010s | 2017-01-21 | Python 2.7.13                                                                |
| Python 3                | 0m0.034s | 0m0.024s | 0m0.007s | 2017-01-21 | Python 3.6.0                                                                 |
| Rust                    | 0m0.002s | 0m0.001s | 0m0.001s | 2017-01-21 | rustc 1.14.0                                                                 |
| Scheme                  | 0m0.005s | 0m0.002s | 0m0.002s | 2017-01-21 | (c) 2008-2016, The CHICKEN Team...Version 4.11.0 (rev ce980c4)...            |
| Tcl                     | 0m0.013s | 0m0.005s | 0m0.006s | 2017-01-21 | Tcl 8.5.9                                                                    |



Running `time ./hello_arg 'Command Line'` five	times for each binary and averaging the results.

| Language | real     | user     | sys      | Date       | Version                                                           |
| -------- | ---:     | ---:     | --:      | :--:       | -------                                                           |
| BASH     | 0m0.005s | 0m0.002s | 0m0.002s | 2017-01-21 | GNU bash, version 4.4.5(1)-release (x86_64-apple-darwin15.6.0)... |
| C        | 0m0.002s | 0m0.001s | 0m0.001s | 2017-01-21 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                    |
| Crystal  | 0m0.004s | 0m0.002s | 0m0.002s | 2017-01-21 | Crystal 0.20.4 (2017-01-06)                                       |
| D        | 0m0.003s | 0m0.001s | 0m0.001s | 2017-01-21 | DMD64 D Compiler v2.072.2                                         |
| Fantom   | 0m0.984s | 0m2.086s | 0m0.156s | 2017-01-21 | ...fan.version: 1.0.69...                                         |
| Go       | 0m0.003s | 0m0.001s | 0m0.002s | 2017-01-21 | go version go1.7.4 darwin/amd64                                   |
| ooc      | 0m0.004s | 0m0.002s | 0m0.002s | 2017-01-21 | rock 0.9.10 codename rita, built on Sun Apr 10 01:50:13 2016      |
| Python 2 | 0m0.019s | 0m0.010s | 0m0.008s | 2017-01-21 | Python 2.7.13                                                     |
| Python 3 | 0m0.035s | 0m0.025s | 0m0.007s | 2017-01-21 | Python 3.6.0                                                      |
| Rust     | 0m0.002s | 0m0.001s | 0m0.001s | 2017-01-21 | rustc 1.14.0                                                      |
| Scheme   | 0m0.005s | 0m0.002s | 0m0.002s | 2017-01-21 | (c) 2008-2016, The CHICKEN Team...Version 4.11.0 (rev ce980c4)... |
| Tcl      | 0m0.012s | 0m0.005s | 0m0.006s | 2017-01-21 | Tcl 8.5.9                                                         |


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

C++
: `clang++ -lstdc++ appname.cpp -o appname`

Crystal
: `crystal build --release appname.cr`

D
: `dmd appname.d`

Go
: `go build appname.go`

ooc
: `rock -v appname.occ`

Rust
: `rustc appname.rs`

Scheme
: `csc appname.scm`


Tool Version
------------

Assember - GAS
: `as --version`
: `gcc --version`

Assember - NASM
: `nasm -v`

BASH
: `bash --version`

C
: `gcc --version`

Crystal
: `crystal version`

D
: `dmd --version`

Fantom
: `fan -version`

GAS
: `as --version`

Go
: `go version`

NASM
: `nasm -v`

ooc
: `rock -V`

Python
: `python --version`

Rust
: `rustc -V`

Scheme
: CHICKEN: `csc -version`

Tcl
: `echo 'puts [info patchlevel]' | tclsh`


ToDo
----

* [x] Setup proper testing `make test`
* [x] Setup make file to build all of these
* [ ] Figure out how to code `hello_arg.asm`
* [ ] Figure out how to code `hello_arg.cpp`





[CHICKEN Scheme]: https://www.call-cc.org/
