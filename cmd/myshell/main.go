package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getBuiltins() []string {
	return []string{"echo", "exit", "type"}
}

func pathCmd(inputTokens []string) (string, bool) {
	cmd := inputTokens[0]
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		cmdPath := fmt.Sprintf("%s/%s", path, cmd)
		if _, err := os.Stat(cmdPath); !errors.Is(err, os.ErrNotExist) {
			return cmdPath, true
		}
	}
	return "", false
}

func execCmd(inputTokens []string) {
	cmd := inputTokens[0]
	switch cmd {
	case "echo":
		fmt.Println(strings.Join(inputTokens[1:], " "))
	case "exit":
		exitCode := 0
		if len(inputTokens) > 1 {
			exitCode, _ = strconv.Atoi(inputTokens[1])
		}

		os.Exit(exitCode)
	case "type":
		typeOption := inputTokens[1]
		if slices.Contains(getBuiltins(), typeOption) {
			fmt.Printf("%s is a shell builtin\n", typeOption)
		} else if cmdPath, found := pathCmd(inputTokens[1:]); found {
			fmt.Printf("%s is %s\n", typeOption, cmdPath)
		} else {
			fmt.Printf("%s: not found\n", typeOption)
		}
	default:
		fmt.Printf("%s: command not found\n", cmd)
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
