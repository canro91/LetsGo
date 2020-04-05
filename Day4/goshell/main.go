package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	splits := strings.Split(input, " ")
	name := splits[0]
	args := splits[1:]

	switch name {
	case "cd":
		if len(args) == 0 {
			homeDir, _ := os.UserHomeDir()
			return os.Chdir(homeDir)
		} else {
			return os.Chdir(args[0])
		}

	case "exit":
		os.Exit(0)
	}

	var command = exec.Command(name, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
