package main

import "fmt"

func main2() {
	fmt.Println("Hello,world1")

	values := []int{3, 2, 5, 6, 1, 9}
	fmt.Println(values)

	quickSort(values)
	fmt.Println(values)
}

func quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	pivot := values[0]
	head, i, tail := 0, 1, len(values)-1
	for head < tail {
		if values[i] > pivot {
			values[tail], values[i] = values[i], values[tail]
			tail--
		} else {
			values[head], values[i] = values[i], values[head]
			head++
			i++
		}
	}
	quickSort(values[:head])
	quickSort(values[head+1:])
}
