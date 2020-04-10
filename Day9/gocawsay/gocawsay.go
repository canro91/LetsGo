package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func convertTabsToSpaces(lines []string) []string {
	var output []string
	for _, line := range lines {
		noTabs := strings.Replace(line, "\t", "    ", -1)
		output = append(output, noTabs)
	}
	return output
}

func calculateMaxWidth(lines []string) int {
	width := 0
	for _, line := range lines {
		l := utf8.RuneCountInString(line)
		if l > width {
			width = l
		}
	}
	return width
}

func normalizeLength(lines []string, width int) []string {
	var output []string
	for _, line := range lines {
		len := utf8.RuneCountInString(line)
		l := line + strings.Repeat(" ", width-len)
		output = append(output, l)
	}
	return output
}

func buildBallon(lines []string, width int) string {
	var output []string
	top := " " + strings.Repeat("_", width+2)
	output = append(output, top)

	count := len(lines)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", "<", lines[0], ">")
		output = append(output, s)
	} else {
		s := fmt.Sprintf("%s %s %s", "/", lines[0], "\\")
		output = append(output, s)

		for i := 1; i < count-1; i++ {
			s = fmt.Sprintf("%s %s %s", "|", lines[i], "|")
			output = append(output, s)
		}

		s = fmt.Sprintf("%s %s %s", "\\", lines[count-1], "/")
		output = append(output, s)
	}

	bottom := " " + strings.Repeat("-", width+2)
	output = append(output, bottom)

	return strings.Join(output, "\n")
}

func printFigure(code string) {
	var cow = `   \  ^__^
    \ (oo)\_______
      (__)\       )\/\
          ||----w |
          ||     ||
	  `

	var bat = `     \
      \
   =,    (\_/)    ,=
    /'-'--(")--'-'\
   /     (___)     \
  /.-.-./ " " \.-.-.\
`

	var cat = `       \
        \
         /\_/\
    ____/ o o \
  /~____  =ø= /
 (______)__m_m)
`

	switch code {
	case "cow":
		fmt.Println(cow)
	case "cat":
		fmt.Println(cat)
	case "bat":
		fmt.Println(bat)
	default:
		fmt.Println("Sorry, figure not supported. Try with cow")
	}
}

func main() {
	if !isInputFromPipe() {
		fmt.Fprintln(os.Stderr, "Use a pipe, please")
	} else {
		reader := bufio.NewReader(os.Stdin)
		// Why not simply:
		//input, _ := reader.ReadString('\n')
		var lines []string
		for {
			input, _, err := reader.ReadLine()
			if err != nil || err == io.EOF {
				break
			}

			lines = append(lines, string(input))
		}

		var figure string
		flag.StringVar(&figure, "f", "cow", "Figure name. Try 'cow' or 'stegosaurus'")
		flag.Parse()

		lines = convertTabsToSpaces(lines)
		maxWidth := calculateMaxWidth(lines)
		messages := normalizeLength(lines, maxWidth)
		ballon := buildBallon(messages, maxWidth)
		fmt.Println(ballon)
		printFigure(figure)
		fmt.Println()
	}
}
