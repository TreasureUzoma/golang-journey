package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
func main() {

	fmt.Print("Enter a number: ");
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	num, err := strconv.ParseFloat(input, 64);

	if err != nil {
		fmt.Println("Invalid input. Please use a valid number.")
		return
	}

	if num < 0 {
		fmt.Println("Cannot calculate square root of a negative nymber")
		return
	}

	sqrtResult := math.Sqrt(num)

	fmt.Printf("The square root of %.2f is %.2f\n", num, sqrtResult)
	
}