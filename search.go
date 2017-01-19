package main

import (
	"fmt"
)

func main() {
	values := []int{1, 3, 5, 6, 7, 9}
	i := binarySearch2(1, values)
	fmt.Println(i)
}

func binarySearch1(someValue int, values []int) int {
	if len(values) == 0 {
		return -1
	}
	head, tail := 0, len(values)-1
	middle := (head + tail) / 2
	if someValue == values[middle] {
		return middle
	} else if someValue > middle {
		//fmt.Println(values[middle:])
		binarySearch1(someValue, values[middle:])
	} else {
		binarySearch1(someValue, values[:middle-1])
	}
	return -1
}
func binarySearch2(someValue int, values []int) int {
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
