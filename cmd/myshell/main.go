package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		inputTokens := strings.Split(input, " ")
		switch inputTokens[0] {
		case "exit":
			exitCode, err := strconv.Atoi(inputTokens[1])
			if err != nil {
				exitCode = 0
			}
			os.Exit(exitCode)
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}

}
