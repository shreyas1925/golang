package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	// var num1 int = 5
	// var num2 float64 = 12.8

	// fmt.Println("Heyy ", num1+int(num2))

	// Two ways:
	// 		math/rand
	// 		crypto/rand

	// Let's try using math/rand where seed is used to set some units like if we
	// set time definetely the value we will be getting will be random

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 2)

	// Let's try using crypto/rand ? But why? Because it is took from the cryptography
	// algorithm where we can belive of better accuracy and perfect randomness'

	RandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(RandomNumber)

}
