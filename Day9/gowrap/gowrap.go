package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

const lineWidth = 50

func main() {
	if !isInputFromPipe() {
		fmt.Fprintln(os.Stderr, "Use a pipe, please!")
	} else {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		var width int
		flag.IntVar(&width, "n", lineWidth, "Line width. By default: 50")
		flag.Parse()

		var output []string
		spaceLeft := width
		line := ""

		words := strings.Fields(input)
		for _, word := range words {
			len := utf8.RuneCountInString(word)
			if len+1 > spaceLeft {
				output = append(output, line)
				line = word + " "
				spaceLeft = width - (len + 1)
			} else {
				line += word + " "
				spaceLeft -= (len + 1)
			}
		}

		if len(line) != 0 {
			output = append(output, line)
		}

		fmt.Println(strings.Join(output, "\n"))
	}
}
