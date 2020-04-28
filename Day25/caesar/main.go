package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

func main() {
	encryptCommand := flag.NewFlagSet("enc", flag.ExitOnError)
	decryptCommand := flag.NewFlagSet("dec", flag.ExitOnError)

	var k int
	var text string
	encryptCommand.IntVar(&k, "k", 3, "Shift value")
	encryptCommand.StringVar(&text, "t", "", "Text to encrypt")

	decryptCommand.IntVar(&k, "k", 3, "Shift value")
	decryptCommand.StringVar(&text, "t", "", "Text to decrypt")

	if len(os.Args) < 2 {
		fmt.Println("encrypt or decrypt subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "e", "enc", "encrypt":
		encryptCommand.Parse(os.Args[2:])
	case "d", "dec", "decrypt":
		decryptCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if encryptCommand.Parsed() {
		if text == "" {
			encryptCommand.PrintDefaults()
			os.Exit(1)
		}
		encrypt(k, text)
	}

	if decryptCommand.Parsed() {
		if text == "" {
			decryptCommand.PrintDefaults()
			os.Exit(1)
		}
		decrypt(k, text)
	}
}

func encrypt(k int, plainText string) {
	shift(plainText, k)
}

func decrypt(k int, cipherText string) {
	shift(cipherText, -k)
}

func shift(text string, shift int) {
	for _, c := range text {
		if !unicode.IsLetter(c) {
			fmt.Print(string(c))
			continue
		}

		var factor int
		if unicode.IsUpper(c) {
			factor = 'A'
		} else {
			factor = 'a'
		}

		e := (((int(c) + shift) - factor) % 26) + factor
		fmt.Print(string(e))
	}
	fmt.Println()
}
