package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, error := bufio.NewReader(os.Stdin).ReadString('\n')
	if error != nil {
		panic("couldnt read from stdin")
	}

	input = strings.TrimSuffix(input, "\n")
	fmt.Printf("%s: command not found\n", input)

}
