package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// fmt.Println("Hello!!")
	performGetReuest()
}
func performGetReuest() {

	const url = "http://localhost:5000/get"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code of our grt response is ", res.StatusCode)
	fmt.Println("Content length :", res.ContentLength)

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Content are : ", string(content))
}
