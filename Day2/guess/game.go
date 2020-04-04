package main

import (
	"time"
	"math/rand"
	"fmt"
)

const maxNumber = 20
const retries = 3

func main(){
		number := rand.Intn(maxNumber)
	won := false

	fmt.Println("Guess the number. It's between 0 and", maxNumber, ". You have", retries, " retries")
	for retry := retries; retry >= 1; retry -- {
		fmt.Print("> ")
		var input int
		fmt.Scanf("%d", &input)
		if input < 0 || input > maxNumber {
			fmt.Println("Dude, did I told you the number is between 0 and,", maxNumber, "?")
			fmt.Println("You have lost one retry...")
		} else if input == number {
			fmt.Println("That's it...You won")
			won = true
			break
		} else if input < number {
			fmt.Println("Your number is smaller than that. Try a greater one. Now you have", retry, "retries")
		} else {
			fmt.Println("You tried too high this time. You have, ", retry, "left")
		}
	}

	if !won {
		fmt.Println("The number was", number)
		fmt.Println("Sorry, you lost this time. Keep trying")
	}
}