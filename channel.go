package main

import (
	"context"
	"fmt"
	"strconv"
)

func main3() {
	//var queue chan int
	queue := make(chan int, 10)
	for i := 0; i < 5; i++ {
		queue <- i
	}
	//i := len(queue)
	for len(queue) > 0 {
		value := <-queue
		fmt.Println(value)
	}
	fmt.Println(len(queue))
	fmt.Println(cap(queue))

	q := NewQ()

	for i := 1; i < 8; i++ {
		t := new(task)
		t.Name = "s" + strconv.FormatInt(int64(i), 10)
		q.Put(t)
	}
	for q.Len() != 0 {
		t := q.Get()
		fmt.Println(t.Name)
	}

}

type task struct {
	Name    string
	Fn      func(...interface{})
	p       int
	context context.Context //withTimeout
}

type queue struct {
	h chan *task
	m chan *task
}

func NewQ() *queue {
	return &queue{h: make(chan *task, 1000), m: make(chan *task, 1000)}
}

func (q *queue) Put(t *task) {
	if t.p == 0 {
		q.h <- t
	} else {
		q.m <- t
	}
}

func (q *queue) Get() *task {
	if len(q.h) > 0 {
		return <-q.h
	}
	return <-q.m
}

func (q *queue) Len() int {
	return len(q.h) + len(q.m)
}
