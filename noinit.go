package main

import (
	"fmt"
)

type price struct {
	value int
}

func mainn() {
	var p price
	p.value = 3
	fmt.Println(p.value)
	fmt.Println(3 << 3)
	//s := "Go编程"
	fmt.Println([]rune("好"))
}
