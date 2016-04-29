Hello Notes
===========

Notes regarding the compilation of simple "Hello, world!" apps.


| Language         | Category                               | Source Size | Binary Bytes | Binary Size | Compile Line                                                                  | Notes                                                 |
| --------         | --------                               | ----------: | -----------: | ----------: | ------------                                                                  | -----                                                 |
| Assembler (GAS)  | static string output                   | 459 B     | 8,288        | 8.1 KB      | `as hello_str_64.s && ld -o hello_str_64 -e _main hello_str_64.o` |                                                       |
| Assembler (NASM) | static string output                   | 1,199 B     | 8,288        | 8.1 KB      | `nasm -f macho hello_str.asm && ld -o hello_str -e mystart hello_str.o`       |                                                       |
| C                | static string output                   | 73 B        | 8,432        | 8.2 KB      | `gcc -ansi hello_str.c -o hello_str`                                          |                                                       |
| C                | variable string output                 | 94 B        | 8,432        | 8.2 KB      | `gcc -ansi hello_var.c -o hello_var`                                          |                                                       |
| C                | argument string output                 | 102 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_arg.c -o hello_arg`                                          | `argv[1]` in the `printf`                             |
| C                | argument variable string output        | 115 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_argvar.c -o hello_argvar`                                    | Pointer variable                                      |
| C                | argument memcpy variable string output | 154 B       | 8,432        | 8.2 KB      | `gcc -ansi hello_argvar_memcpy.c -o hello_argvar_memcpy`                      |                                                       |
| C                | argument strcpy variable string output | 150 B       | 8,496        | 8.3 KB      | `gcc -ansi hello_argvar_strcpy.c -o hello_argvar_strcpy`                      |                                                       |
| Rust             | static string output                   | 45 B        | 302,620      | 303 KB      | `rustc hello_var.rs`                                                          |                                                       |
| Rust             | variable string output                 | 63 B        | 302,860      | 303 KB      | `rustc hello_var.rs`                                                          |                                                       |
| Rust             | arg string output                      | 127 B       | 334,420      | 336 KB      | `rustc hello_arg.rs`                                                          |                                                       |
| Go               | minimalist static string output        | 55 B        | 1,070,112    | 1.1 MB      | `go build hello_min.go`                                                       | Didn't import the `fmt` package. Just used `println`. |
| Go               | static string output                   | 73 B        | 2,253,776    | 2.3 MB      | `go build hello_str.go`                                                       |                                                       |
| Go               | variable string output                 | 88 B        | 2,253,760    | 2.3 MB      | `go build hello_var.go`                                                       | `fmt.Printf` is smaller than `fmt.Println`?           |
| Go               | arg string output                      | 94 B        | 2,253,760    | 2.3 MB      | `go build hello_arg.go`                                                       | `os.Args` package is free with `fmt` package?         |


Compile Line
------------

Assembler (NASM)
: `

C
: `gcc -ansi source.c -o binary_name`

Go
: `go build source.go`

Rust
: `rustc source.rs`

