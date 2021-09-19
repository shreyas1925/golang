package main

import (
	"fmt"
	"time"
)

func main() {

	currentDataTime := time.Now()

	fmt.Println(currentDataTime.Format("01-02-2006 15:04:00 Monday"))

	createdateTime := time.Date(2001, time.April, 19, 12, 14, 19, 25, time.UTC)

	fmt.Println(createdateTime.Format("01-02-2006 15:04:00 Monday"))
}
