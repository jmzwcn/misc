package main

import (
	"fmt"
)

func main() {
	//c := make(chan bool)
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	//<-c
}
