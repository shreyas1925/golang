package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Email address:")
	name := bufio.NewReader(os.Stdin)
	inputname, _ := name.ReadString('\n')

	fmt.Println("Password:")
	password := bufio.NewReader(os.Stdin)
	inputpassword, _ := password.ReadString('\n')

	fmt.Println("Your credentials are: ")
	fmt.Println(inputname, inputpassword)
	fmt.Println("Registered successfully!!")

}
