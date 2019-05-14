Hello Notes
===========

Notes regarding the compilation of simple "Hello, world!" apps.

Compiled Languages
------------------

| Language                      | Category                           | Source Size | Binary Bytes | Binary Size | Notes                                                 |
| --------                      | --------                           | ----------: | -----------: | ----------: | -----                                                 |
| Assembler (GAS - LLVM v8.0.0) | static string output               | 458 B       | 8,328        | 8.1 KB      |                                                       |
| Assembler (NASM v0.98.40)     | static string output               | 1,199 B     | 8,288        | 8.1 KB      |                                                       |
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
| Go v1.7.4                     | minimalist static string output    | 55 B        | 971,584      | 948.81 KB   | Didn't import the `fmt` package. Just used `println`. |
| Go v1.7.4                     | static string output               | 73 B        | 1,611,712    | 1.54 MB     |                                                       |
| Go v1.7.4                     | variable string output             | 89 B        | 1,620,096    | 1.55 MB     | `fmt.Printf` is smaller than `fmt.Println`?           |
| Go v1.7.4                     | arg variable string output         | 122 B       | 1,620,128    | 1.55 MB     | `os.Args` package is free with `fmt` package?         |
| Go v1.7.4                     | environment variable string output | 130 B       | 1,624,256    | 1.55 MB     |                                                       |
| ooc (rock 0.9.10)             | static string output               | 25 B        | 612,580      | 598.22 KB   |                                                       |
| ooc (rock 0.9.10)             | variable string output             | 37 B        | 612,668      | 598.31 KB   |                                                       |
| ooc (rock 0.9.10)             | arg variable string output         | 72 B        | 616,820      | 602.36 KB   |                                                       |
| ooc (rock 0.9.10)             | environment variable string output | 62 B        | 616,924      | 602.47 KB   |                                                       |
| Rust (rustc 1.14.0)           | static string output               | 44 B        | 346,504      | 338.38 KB   |                                                       |
| Rust (rustc 1.14.0)           | variable string output             | 62 B        | 346,696      | 338.57 KB   |                                                       |
| Rust (rustc 1.14.0)           | arg variable string output         | 123 B       | 379,784      | 370.88 KB   |                                                       |
| Rust (rustc 1.14.0)           | environment variable string output | 134 B       | 355,088      | 346.77 KB   |                                                       |
| Scheme (Chicken 4.11.0)       | static string output               | 34 B        | 14,160       | 13.83 KB    | Compiled using [CHICKEN Scheme][]                     |
| Scheme (Chicken 4.11.0)       | variable string output             | 74 B        | 14,272       | 13.94 KB    | Compiled using [CHICKEN Scheme][]                     |
| Scheme (Chicken 4.11.0)       | arg variable string output         | 84 B        | 14,280       | 13.95 KB    | Compiled using [CHICKEN Scheme][]                     |
| Scheme (Chicken 4.11.0)       | environment variable string output | 88 B        | 14,232       | 13.90 KB    | Compiled using [CHICKEN Scheme][]                     |


| Str Bytes | Language         |
| --------: | --------         |
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

| Language  | Simplicity | Speed | Binary Size | Total | Simplicity & Speed |
| --------  | :--------: | :---: | :---------: | :---: | :----------------: |
| Assembler | 5          | 1     | 2           | 8     | 6                  |
| C         | 4          | 1     | 2           | 7     | 5                  |
| C++       | 3          | 1     | 3           | 7     | 4                  |
| Crystal   | 2          | 1     | 4           | 7     | 3                  |
| D         | 2          | 1     | 4           | 8     | 3                  |
| Go        | 3          | 1     | 5           | 9     | 4                  |
| ooc       | 1          | 2     | 4           | 7     | 3                  |
| Rust      | 3          | 1     | 4           | 8     | 4                  |
| Scheme    | 2          | 1     | 3           | 6     | 3                  |

<small>Lower numbers are better.</small>

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


### Speed Runs

I realize these are crap example of testing speed. The applications are too simple. But it still gives some reference for discussion.

Running `time ./hello_str` five times for each binary, removing the max and min values, then averaging the the three that are left. Times are in seconds.

| Language                | real    | user    | sys     | Date       | Version                                                                      |
| --------                | ---:    | ---:    | --:     | :--:       | -------                                                                      |
| Assembler (GAS 64-bit)  | 0.00300 | 0.00100 | 0.00100 | 2017-01-22 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                               |
| Assembler (NASM 32-bit) | 0.00200 | 0.00000 | 0.00000 | 2017-01-22 | NASM version 0.98.40 (Apple Computer, Inc. build 11) compiled on Nov 15 2016 |
| C                       | 0.00200 | 0.00100 | 0.00100 | 2017-01-22 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                               |
| C++                     | 0.00300 | 0.00100 | 0.00100 | 2017-01-22 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                               |
| Crystal                 | 0.00433 | 0.00200 | 0.00200 | 2017-01-22 | Crystal 0.20.4 (2017-01-06)                                                  |
| D                       | 0.00300 | 0.00100 | 0.00100 | 2017-01-22 | DMD64 D Compiler v2.072.2                                                    |
| Go                      | 0.00367 | 0.00000 | 0.00000 | 2017-01-22 | go version go1.7.4 darwin/amd64                                              |
| ooc                     | 0.00400 | 0.00200 | 0.00200 | 2017-01-22 | rock 0.9.10 codename rita, built on Sun Apr 10 01:50:13 2016                 |
| Rust                    | 0.00300 | 0.00100 | 0.00100 | 2017-01-22 | rustc 1.14.0                                                                 |
| Scheme                  | 0.00500 | 0.00200 | 0.00200 | 2017-01-22 | (c) 2008-2016, The CHICKEN Team...Version 4.11.0 (rev ce980c4)...            |



Running `time ./hello_arg 'Command Line'` five times for each binary, removing the max and min values, then averaging the the three that are left. Times are in seconds.

| Language | real    | user    | sys     | Date       | Version                                                           |
| -------- | ---:    | ---:    | --:     | :--:       | -------                                                           |
| C        | 0.00567 | 0.00267 | 0.00267 | 2017-01-22 | Apple LLVM version 8.0.0 (clang-800.0.42.1)...                    |
| Crystal  | 0.01100 | 0.00533 | 0.00533 | 2017-01-22 | Crystal 0.20.4 (2017-01-06)                                       |
| D        | 0.00900 | 0.00267 | 0.00267 | 2017-01-22 | DMD64 D Compiler v2.072.2                                         |
| Go       | 0.01000 | 0.00133 | 0.00133 | 2017-01-22 | go version go1.7.4 darwin/amd64                                   |
| ooc      | 0.01067 | 0.00533 | 0.00533 | 2017-01-22 | rock 0.9.10 codename rita, built on Sun Apr 10 01:50:13 2016      |
| Rust     | 0.00800 | 0.00267 | 0.00267 | 2017-01-22 | rustc 1.14.0                                                      |
| Scheme   | 0.01333 | 0.00533 | 0.00533 | 2017-01-22 | (c) 2008-2016, The CHICKEN Team...Version 4.11.0 (rev ce980c4)... |


### Compile Line

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


Scripting Languages
-------------------


| Language          | Category                           | Source Size | Binary Size         | Package Size             | Interpreter                                                   |
| --------          | --------                           | ----------: | ----------:         | -----------:             | -----------                                                   |
| BASH v4.4.5       | static string output               | 41 B        | 946.56 KB           | (source) 8.94 MB         | `bash` shell                                                  |
| BASH v4.4.5       | variable string output             | 49 B        | 946.56 KB           | (source) 8.94 MB         | `bash` shell                                                  |
| BASH v4.4.5       | argument string output             | 48 B        | 946.56 KB           | (source) 8.94 MB         | `bash` shell                                                  |
| BASH v4.4.5       | environment variable string output | 52 B        | 946.56 KB           | (source) 8.94 MB         | `bash` shell                                                  |
| Fantom v1.0.69    | static string output               | 84 B        | 2.27 KB             | 13.16 MB                 | `fan` interpreter (BASH script that launches Java Jar file)   |
| Fantom v1.0.69    | variable string output             | 98 B        | 2.27 KB             | 13.16 MB                 | `fan` interpreter (BASH script that launches Java Jar file)   |
| Fantom v1.0.69    | arg variable string output         | 129 B       | 2.27 KB             | 13.16 MB                 | `fan` interpreter (BASH script that launches Java Jar file)   |
| Fantom v1.0.69    | environment variable string output | 133 B       | 2.27 KB             | 13.16 MB                 | `fan` interpreter (BASH script that launches Java Jar file)   |
| JSDB 1.8.0.7      | static string output               | 45 B        | 3.49 MB             | 0.99 MB                  | `jsdb` interpreter (SpiderMonkey based binary)                |
| JSDB 1.8.0.7      | variable string output             | 56 B        | 3.49 MB             | 0.99 MB                  | `jsdb` interpreter (SpiderMonkey based binary)                |
| JSDB 1.8.0.7      | arg variable string output         | 74 B        | 3.49 MB             | 0.99 MB                  | `jsdb` interpreter (SpiderMonkey based binary)                |
| JSDB 1.8.0.7      | environment variable string output | 82 B        | 3.49 MB             | 0.99 MB                  | `jsdb` interpreter (SpiderMonkey based binary)                |
| Lua 5.2.4         | static string output               | 42 B        | 19.30 KB            | 173.38 KB                | `lua` interpreter                                             |
| Lua 5.2.4         | variable string output             | 49 B        | 19.30 KB            | 173.38 KB                | `lua` interpreter                                             |
| Lua 5.2.4         | arg variable string output         | 60 B        | 19.30 KB            | 173.38 KB                | `lua` interpreter                                             |
| Lua 5.2.4         | environment variable string output | 72 B        | 19.30 KB            | 173.38 KB                | `lua` interpreter                                             |
| Nashorn 1.8.0_111 | static string output               | 41 B        | 57.11 KB            | (JDK) 174.76 MB          | `jjs` interpreter                                             |
| Nashorn 1.8.0_111 | variable string output             | 52 B        | 57.11 KB            | (JDK) 174.76 MB          | `jjs` interpreter                                             |
| Nashorn 1.8.0_111 | arg variable string output         | 63 B        | 57.11 KB            | (JDK) 174.76 MB          | `jjs` interpreter                                             |
| Nashorn 1.8.0_111 | environment variable string output | 65 B        | 57.11 KB            | (JDK) 174.76 MB          | `jjs` interpreter                                             |
| Node v7.4.0       | static string output               | 49 B        | 31.43 MB            | 9.84 MB                  | `node` interpreter (V8 based binary)                          |
| Node v7.4.0       | variable string output             | 60 B        | 31.43 MB            | 9.84 MB                  | `node` interpreter (V8 based binary)                          |
| Node v7.4.0       | arg variable string output         | 85 B        | 31.43 MB            | 9.84 MB                  | `node` interpreter (V8 based binary)                          |
| Node v7.4.0       | environment variable string output | 87 B        | 31.43 MB            | 9.84 MB                  | `node` interpreter (V8 based binary)                          |
| Python 2.7.13     | static string output               | 44 B        | 13.00 KB            | 16.29 MB                 | `python` interpreter                                          |
| Python 2.7.13     | variable string output             | 58 B        | 13.00 KB            | 16.29 MB                 | `python` interpreter                                          |
| Python 2.7.13     | arg variable string output         | 74 B        | 13.00 KB            | 16.29 MB                 | `python` interpreter                                          |
| Python 2.7.13     | environment variable string output | 81 B        | 13.00 KB            | 16.29 MB                 | `python` interpreter                                          |
| Python 3.6.0      | static string output               | 46 B        | 13.30 KB            | 21.23 MB                 | `python3` interpreter                                         |
| Python 3.6.0      | variable string output             | 66 B        | 13.30 KB            | 21.23 MB                 | `python3` interpreter                                         |
| Python 3.6.0      | arg variable string output         | 83 B        | 13.30 KB            | 21.23 MB                 | `python3` interpreter                                         |
| Python 3.6.0      | environment variable string output | 89 B        | 13.30 KB            | 21.23 MB                 | `python3` interpreter                                         |
| REXX 5.00         | static string output               | 40 B        | 348.90 KB           | 2.73 MB                  | `rexx` interpreter (Regina 3.9.1)                             |
| REXX 5.00         | variable string output             | 47 B        | 348.90 KB           | 2.73 MB                  | `rexx` interpreter (Regina 3.9.1)                             |
| REXX 5.00         | arg variable string output         | 50 B        | 348.90 KB           | 2.73 MB                  | `rexx` interpreter (Regina 3.9.1)                             |
| REXX 5.00         | environment variable string output | 73 B        | 348.90 KB           | 2.73 MB                  | `rexx` interpreter (Regina 3.9.1)                             |
| RingoJS 1.0.0     | static string output               | 50 B        | (`run.jar`) 4.97 KB | 4.32 MB + JRE = 74.58 MB | `ringo` interpreter (BASH script that launches Java Jar file) |
| RingoJS 1.0.0     | variable string output             | 61 B        | (`run.jar`) 4.97 KB | 4.32 MB + JRE = 74.58 MB | `ringo` interpreter (BASH script that launches Java Jar file) |
| RingoJS 1.0.0     | arg variable string output         | 129 B       | (`run.jar`) 4.97 KB | 4.32 MB + JRE = 74.58 MB | `ringo` interpreter (BASH script that launches Java Jar file) |
| RingoJS 1.0.0     | environment variable string output | 162 B       | (`run.jar`) 4.97 KB | 4.32 MB + JRE = 74.58 MB | `ringo` interpreter (BASH script that launches Java Jar file) |
| Tcl 8.5.9         | static string output               | 42 B        | 49.41 KB            | 21.90 MB                 | `tclsh` shell                                                 |
| Tcl 8.5.9         | variable string output             | 52 B        | 49.41 KB            | 21.90 MB                 | `tclsh` shell                                                 |
| Tcl 8.5.9         | arg variable string output         | 53 B        | 49.41 KB            | 21.90 MB                 | `tclsh` shell                                                 |
| Tcl 8.5.9         | environment variable string output | 48 B        | 49.41 KB            | 21.90 MB                 | `tclsh` shell                                                 |


Binary Size
: The size of the main executable of the shell or interpretter

Package Size
: The size of the Linux Tar.GZip or Zip archive. This was the most consistent downloadable install package format.


| Str Bytes | Language         |
| --------: | --------         |
| `0000041` | BASH             |
| `0000042` | Tcl              |
| `0000044` | Python 2         |
| `0000058` | Python 3         |
| `0000084` | Fantom           |



### Basic Finding

| Language | Simplicity | Speed | Binary Size | Total | Simplicity & Speed |
| -------- | :--------: | :---: | :---------: | :---: | :----------------: |
| BASH     | 1          | 1     | 1           | 3     | 6                  |
| Fantom   | 3          | 4     | 1           | 8     | 6                  |
| Node     | 1          |       |             |       |                    |
| Python 2 | 2          | 2     | 1           | 4     | 3                  |
| Python 3 | 2          | 2     | 1           | 4     | 3                  |
| RingoJS  | 1          |       |             |       |                    |
| Tcl      | 1          | 2     | 1           | 3     | 3                  |

<small>Lower numbers are better.</small>

| Rank | Language Simplicity               | Binary Size | Binary Speed    |
| :--: | :-----------------:               | :---------: | :----------:    |
| 1    | Very simple for most basic tasks  | < 1 KB      | < 0m0.003s real |
| 2    | Simple for most basic tasks       | < 10 KB     | < 0m0.005s real |
| 3    | OK for most basic tasks           | < 100 KB    | < 0m0.007s real |
| 4    | Complex for most basic tasks      | < 1 MB      | < 0m0.009s real |
| 5    | Very complex for most basic tasks | >= 1 MB     | > 0m0.008s real |


### Speed Runs

Running `time ./hello_str.ext` five times for each script, removing the max and min values, then averaging the the three that are left. Times are in seconds.

| Language | real    | user    | sys     | Date       | Version                                                           |
| -------- | ---:    | ---:    | --:     | :--:       | -------                                                           |
| BASH     | 0.00400 | 0.00100 | 0.00100 | 2017-01-22 | GNU bash, version 4.4.5(1)-release (x86_64-apple-darwin15.6.0)... |
| Fantom   | 0.81133 | 2.00233 | 2.00233 | 2017-01-22 | ...fan.version: 1.0.69...                                         |
| Lua      | 0.00400 | 0.00100 | 0.00100 | 2017-01-23 |                                                                   |
| REXX     | 0.00400 | 0.00100 | 0.00100 | 2017-01-22 |                                                                   |
| Python 2 | 0.01800 | 0.01000 | 0.01000 | 2017-01-22 | Python 2.7.13                                                     |
| Python 3 | 0.03200 | 0.02400 | 0.02400 | 2017-01-22 | Python 3.6.0                                                      |
| Tcl      | 0.01200 | 0.00500 | 0.00500 | 2017-01-22 | Tcl 8.5.9                                                         |



Running `time ./hello_arg.ext 'Command Line'` five times for each script, removing the max and min values, then averaging the the three that are left. Times are in seconds.

| Language | real    | user    | sys     | Date       | Version                                                           |
| -------- | ---:    | ---:    | --:     | :--:       | -------                                                           |
| BASH     | 0.01100 | 0.00300 | 0.00300 | 2017-01-22 | GNU bash, version 4.4.5(1)-release (x86_64-apple-darwin15.6.0)... |
| Fantom   | 2.46533 | 5.67833 | 5.67833 | 2017-01-22 | ...fan.version: 1.0.69...                                         |
| Lua      | 0.01067 | 0.00267 | 0.00267 | 2017-01-23 |                                                                   |
| REXX     | 0.01067 | 0.00267 | 0.00267 | 2017-01-22 |                                                                   |
| Python 2 | 0.04833 | 0.02667 | 0.02667 | 2017-01-22 | Python 2.7.13                                                     |
| Python 3 | 0.08600 | 0.06400 | 0.06400 | 2017-01-22 | Python 3.6.0                                                      |
| Tcl      | 0.03233 | 0.01333 | 0.01333 | 2017-01-22 | Tcl 8.5.9                                                         |



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


Links
-----

* **JavaScript**
	* [JSDB: JavaScript for databases][]
	* [Node.js][]
	* [OpenJDK: Nashorn][]
		* [Nashorn - OpenJDK Wiki][]
	* [RingoJS - Multi-threaded JavaScript on the JVM][]



ToDo
----

* [x] Setup proper testing `make test`
* [x] Setup make file to build all of these
* [ ] Figure out how to code `hello_arg.asm`
* [ ] Figure out how to code `hello_arg.cpp`





[CHICKEN Scheme]: https://www.call-cc.org/
[JSDB: JavaScript for databases]: http://www.jsdb.org/
[Node.js]: https://nodejs.org/
[RingoJS - Multi-threaded JavaScript on the JVM]: https://ringojs.org/
[OpenJDK: Nashorn]: http://openjdk.java.net/projects/nashorn/
[Nashorn - OpenJDK Wiki]: https://wiki.openjdk.java.net/display/Nashorn/Main


