package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Super simple all-in-one
	// if b, err := ioutil.ReadAll(os.Stdin); err == nil {
	// 	fmt.Printf("Hello, %s!\n", string(b))
	// }

	// Really simple line-by-line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("Hello, %s!\n", scanner.Text())
	}

	// Smarter all-in-one
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// Data is from a Pipe
		if b, err := ioutil.ReadAll(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		} else {
			fmt.Printf("Hello, %s!\n", string(b))
		}
	}
}
