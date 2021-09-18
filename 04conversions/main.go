package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter yout rollNo: ")
	reader := bufio.NewReader(os.Stdin)
	rollNo, _ := reader.ReadString('\n')

	fmt.Println("Your roll number is: ", rollNo)

	updatedroll, err := strconv.ParseFloat(strings.TrimSpace(rollNo), 64)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Your updated roll number is: ", updatedroll+5)
	}

}
