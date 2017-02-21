package main

import (
	"fmt"
	"strconv"
)

var steps int

func maingame() {
	array := []int{1, 2, -3, 4, 5}
	fmt.Println(array)
	move(0, array[0], array)
	fmt.Println(steps)
}

func move(index, value int, array []int) {
	newIndex := index + value
	for 0 <= newIndex && newIndex < len(array) {

	}
	if (newIndex) < 0 || (newIndex) > len(array) {
		fmt.Println("will go to " + strconv.Itoa(newIndex))
	} else {
		move(newIndex, array[newIndex], array)
	}
	steps++
}
