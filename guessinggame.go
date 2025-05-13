package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// generate a secret number
	secretNumber := rng.Intn(50) + 1 // generates a number between 0-49
	fmt.Println("I've picked a secret number between 1 and 50.")
	fmt.Println("Try to guess it :)")

	reader := bufio.NewReader(os.Stdin)

	for { // loop indefinitely until the user guesses correctly
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Oops! That doesn't look like a number. Try again (:")
			continue
		}

		// compare guess with secret number
		if guess < secretNumber {
			fmt.Println("Too low!")
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Yay! You guessed it!")
			break // Exit the loop
		}
	}
}
