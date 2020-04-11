package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"math"
)

type RGB struct {
	R, G, B int
}

func toRGB(i int) (color RGB) {
	var f = 0.1
	r := int(math.Sin(f*float64(i)+0)*127 + 128)
	g := int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
	return RGB{r, g, b}
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func main() {
	if !isInputFromPipe() {
		fmt.Fprintln(os.Stderr, "Use a pipe, please")
	} else {
		reader := bufio.NewReader(os.Stdin)
		i := 0
		for {
			input, _, err := reader.ReadRune()
			if err != nil || err == io.EOF {
				break
			}

			color := toRGB(i)
			fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", color.R, color.G, color.B, input)
			i++
		}
		fmt.Println()
	}
}
