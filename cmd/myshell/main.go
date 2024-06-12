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
	return []string{"echo", "exit", "type", "pwd", "cd"}
}

func execCmd(inputTokens []string) {
	cmd := inputTokens[0]
	arguments := inputTokens[1:]
	switch cmd {
	case "echo":
		fmt.Println(strings.Join(arguments, " "))
	case "exit":
		exitCode := 0
		if len(arguments) > 0 {
			exitCode, _ = strconv.Atoi(arguments[0])
		}

		os.Exit(exitCode)
	case "type":
		typeOption := arguments[0]
		if slices.Contains(getBuiltins(), typeOption) {
			fmt.Printf("%s is a shell builtin\n", typeOption)
		} else if cmdPath, err := exec.LookPath(arguments[0]); err == nil {
			fmt.Printf("%s is %s\n", typeOption, cmdPath)
		} else {
			fmt.Printf("%s: not found\n", typeOption)
		}
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
	case "cd":
		err := os.Chdir(arguments[0])
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", arguments[0])
		}
	default:
		path, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Printf("%s: command not found\n", cmd)
		} else {
			out, _ := exec.Command(path, arguments...).CombinedOutput()
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
