package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?vedio=golang&v=cl7_ouTMFh0"

func main() {

	// fmt.Println("Handling urls in golang")

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())

	//Query methods keeps all the queries in a dictionary (well structered manner)
	queryparams := result.Query()
	fmt.Printf("%T\n", queryparams)

	fmt.Println("We can get the results by using key value pairs ", queryparams["vedio"])

	fmt.Println("Params are as follows")

	for _, data := range queryparams {
		fmt.Println(data)
	}

	// We have constructed the  result out of url responses now we can do vice versa

	constructUrl := &url.URL{

		Scheme:  "https",
		Host:    "amazon.com",
		Path:    "/mens/shirts",
		RawPath: "quantity=4",
	}

	resultedUrl := constructUrl.String()
	fmt.Println(resultedUrl)

}
