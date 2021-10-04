package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	// fmt.Println("Learn web request in golang")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of the response is %T\n", res)
	// output :- Type of the response is *http.Response we are getting its reference which we can manipulate

	defer res.Body.Close()

	// defer is like it validates at a moment but it exceutes after all the methods or contents return their outputs

	datatypes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(datatypes)
	// converting http response data into a string format inorder to print it

	fmt.Println(content)

}
