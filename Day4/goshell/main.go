package main

import (
	"os/user"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		user, _ := user.Current()
		hostname, _ := os.Hostname()

		fmt.Printf("%s at %s in %s > ", user.Username, hostname, currentDir())
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

func currentDir() string {
	cwd, _ := os.Getwd()
	homeDir, _ := os.UserHomeDir()

	dir := strings.Replace(cwd, homeDir, "~", 1)
	folders := strings.Split(dir, "/")
	if len(folders) <= 2 {
		return dir
	} else {
		lastTwo := folders[len(folders) - 2:]
		shortCwd := strings.Join(lastTwo, "/")
		return ".../" + shortCwd
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
