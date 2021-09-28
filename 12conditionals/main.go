package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Welcome to if else")

	secretNumber := 19
	var luckyNumber int

	// number := bufio.NewReader(os.Stdin)

	// luckyNumber, _ := number.ReadString('\n')

	fmt.Println("Enter your lucky number : ")
	fmt.Scanf("%d", &luckyNumber)

	if luckyNumber > secretNumber {

		fmt.Println("You are too high 😁")

	} else if luckyNumber < secretNumber {
		fmt.Println("Grow more jk!! 😉")
	} else {
		fmt.Println("🚀 Hurray!! you guessed it right ")
	}

}
