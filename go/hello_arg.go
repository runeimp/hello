package main

import (
	"fmt"
	"os"
)

func main() {
	x := fmt.Sprintf("Hello, %s!", os.Args[1])

	fmt.Printf("%s\n", x)
}