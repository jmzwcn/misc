package main

import (
	"fmt"
)

func mainsearch() {
	values := []int{1, 3, 5, 6, 7, 9}
	i := binarySearch(6, values)
	fmt.Println(i)
}

func binarySearch(someValue int, values []int) int {
	head, tail := 0, len(values)-1
	for head <= tail {
		middle := (head + tail) / 2
		if someValue == values[middle] {
			return middle
		} else if someValue > values[middle] {
			head = middle + 1
		} else {
			tail = middle - 1
		}
	}
	return -1
}
