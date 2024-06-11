package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func execCmd(inputTokens []string) {
	switch inputTokens[0] {
	case "echo":
		fmt.Println(strings.Join(inputTokens[1:], " "))
	case "exit":
		exitCode := 0
		if len(inputTokens) > 1 {
			exitCode, _ = strconv.Atoi(inputTokens[1])
			// if err != nil {
			// 	exitCode = 0
			// }
		}

		os.Exit(exitCode)
	default:
		fmt.Printf("%s: command not found\n", inputTokens[0])
	}
}

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
		execCmd(inputTokens)
	}

}
