package main

import "fmt"

func main() {
	// fmt.Println("Lets see maps in golang!!")


	socials := make(map[string]string)

	socials["github"] = "@shreyas1925"
	socials["twitter"] = "@shreyas__19"
	socials["instagram"] = "@shreyas__19"

	fmt.Println("My favorite socials links: ",socials)

	for key, value := range socials {

		fmt.Printf("The link for %v is %v\n",key,value)
	}

	students := make(map[int]string)

	students[2] = "shreyas"
	students[24] = "john"
	students[19] = "swayan"

	fmt.Println("Students lists are as follows: ",students)

}