package main

import (
	"fmt"
	"os"
)

func main() {
	x := fmt.Sprintf("Hello, %s!", os.Getenv("HELLO"))

	fmt.Printf("%s\n", x)
}