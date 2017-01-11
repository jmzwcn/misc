package main

import "fmt"

func mains() {
	//保存需要排序的Slice
	arr := []int{9, 3, 4, 7, 2, 1, 0, 11, 12, 11, 13, 4, 7, 2, 1, 0, 11, 12}
	//实际用于排序的Slice
	list := make([]int, len(arr))

	copy(list, arr)
	bubbleSort(list)
	fmt.Println("冒泡排序：\t", list)

	copy(list, arr)
	selectSort(list)
	fmt.Println("选择排序：\t", list)

	copy(list, arr)
	quickSort(list)
	fmt.Println("快速排序：\t", list)

	copy(list, arr)
	insertSort(list)
	fmt.Println("插入排序：\t", list)

	copy(list, arr)
	heapSort(list)
	fmt.Println("大堆排序：\t", list)

	copy(list, arr)
	nl := mergeSort(list)
	fmt.Println("归并排序：\t", nl)
}

func bubbleSort(values []int) {
	//head, i, tail := 0, 1, len(values)-1
	for i := 0; i < len(values); i++ {
		for j := 1; j < len(values)-i; j++ {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
			}
		}

	}
}

func selectSort(values []int) {
	for i := 0; i < len(values); i++ {
		maxIndex := 0
		for j := 0; j < len(values)-i; j++ {
			if values[j] > values[maxIndex] {
				maxIndex = j
			}
		}
		values[len(values)-1-i], values[maxIndex] = values[maxIndex], values[len(values)-1-i]
	}
}

func quickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	pivot := values[0]
	head, tail := 0, len(values)-1
	for head < tail {
		if values[head+1] > pivot {
			values[tail], values[head+1] = values[head+1], values[tail]
			tail--
		} else {
			values[head+1], values[head] = values[head], values[head+1]
			head++
		}
	}
	quickSort(values[:head])
	quickSort(values[head+1:])
}

func insertSort(values []int) {
	for i := 1; i < len(values); i++ {
		temp := values[i]
		j := i - 1
		for j >= 0 && temp < values[j] {
			values[j+1] = values[j] //后挪一个
			j--
		}
		values[j+1] = temp
	}
}

func heapSort(values []int) {
	for i := 0; i < len(values); i++ {
		buildMaxHeap(values, len(values)-i)
		values[0], values[len(values)-1-i] = values[len(values)-1-i], values[0]
	}
}
func buildMaxHeap(values []int, size int) {
	for j := size - 1; j >= 0; j-- {
		parent := (j - 1) / 2
		if values[j] > values[parent] {
			values[j], values[parent] = values[parent], values[j]
		}
	}
}

func mergeSort(values []int) []int {
	if len(values) == 1 {
		return values
	}
	middle := len(values) / 2
	left := mergeSort(values[:middle])
	right := mergeSort(values[middle:])
	return merge(left, right)
}

func merge(values1, values2 []int) []int {
	var newValues []int
	i, j := 0, 0
	for i < len(values1) && j < len(values2) {
		if values1[i] < values2[j] {
			newValues = append(newValues, values1[i])
			i++
		} else {
			newValues = append(newValues, values2[j])
			j++
		}
	}
	newValues = append(newValues, values1[i:]...)
	newValues = append(newValues, values2[j:]...)
	return newValues
}
