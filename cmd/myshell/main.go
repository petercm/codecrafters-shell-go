package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, error := reader.ReadString('\n')
		if error != nil {
			panic("couldnt read from stdin")
		}
		input = strings.TrimSuffix(input, "\n")
		fmt.Printf("%s: command not found\n", input)
	}

}
