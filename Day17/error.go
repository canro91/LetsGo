package main

import (
	"errors"
	"fmt"
)

type TooSmallError struct {
	Number int
	Err error
}

func (e *TooSmallError) Error() string {
	return string(e.Number) + ": Too small"
}

func ProcessEven(value int) (int, error) {
	if value == 42 {
		return -1, errors.New("That's too complex. It's the sense of the universe")
	}

	return 2 * value, nil
}

func ProcessOdd(value int) (int, error){
	if value < 5 {
		return -1, &TooSmallError{Number:value}
	}
	return value - 1, nil
}

func Process(value int) (int, error) {
	if value % 2 == 0 {
		// Not recommended
		// Return ProcessEven directly
		p, err := ProcessEven(value)
		if err != nil {
			return -1, err
		}
		return p, nil
	} else {
		return ProcessOdd(value)
	}
}


func main() {

	fmt.Print("Enter an number, please: ")
	var input int
	fmt.Scanf("%d", &input)

	if value, err := Process(input); err != nil {
		if tooSmall, ok := err.(*TooSmallError); ok {
			fmt.Printf("We found a value too small %d\n", tooSmall.Number)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Here's your number back %d\n", value)
	}
}
