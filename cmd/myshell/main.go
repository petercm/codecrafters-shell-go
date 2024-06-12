package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

func getBuiltins() []string {
	return []string{"echo", "exit", "type", "pwd"}
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
		} else if cmdPath, err := exec.LookPath(inputTokens[1]); err == nil {
			fmt.Printf("%s is %s\n", typeOption, cmdPath)
		} else {
			fmt.Printf("%s: not found\n", typeOption)
		}
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
	default:
		path, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Printf("%s: command not found\n", cmd)
		} else {
			out, _ := exec.Command(path, inputTokens[1:]...).CombinedOutput()
			fmt.Printf("%s", out)
		}
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
