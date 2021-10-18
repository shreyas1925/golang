package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	// fmt.Println("Heyy POST server"
	makeJsonPostRequest()
}

func makeJsonPostRequest() {

	const url = "http://localhost:5000/post"

	reqBody := strings.NewReader(`
			{
				"name" : "Adam Jacob",
				"age"  : 21,
				"gender" : "male"
			}
	`)

	res, err := http.Post(url, "application/json", reqBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))

}
