package main

import (
	"fmt"
)

func mains() {
	//c := make(chan bool)
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	//<-c
}
