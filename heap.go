package main

import (
	"fmt"
)

type Node struct {
	Value    string
	Priority int
}

type Heap struct {
	Nodes []*Node
}

func mainh() {
	heap := Heap{}
	for i := 100; i >= 0; i-- {
		node := &Node{Value: "test", Priority: i}
		heap.push(node)
	}
	fmt.Println(heap.poll())
	fmt.Println(heap.peak().Priority)
	fmt.Println(heap.poll())
	fmt.Println(heap.peak().Priority)
}

func (h *Heap) peak() *Node { //just get, not remove it
	temp := h.Nodes[0]
	return temp
}

func (h *Heap) poll() *Node { //get and remove it
	temp := h.Nodes[0]
	h.Nodes = h.Nodes[1 : len(h.Nodes)-1]
	h.sortHeap(h.Nodes)
	return temp
}

func (h *Heap) push(node *Node) {
	h.Nodes = append(h.Nodes, node)
	h.sortHeap(h.Nodes)
}

func (h *Heap) sortHeap(nodes []*Node) {
	for i := 0; i < len(nodes); i++ {
		//buildMinHeap(nodes, i)
		for j := len(nodes) - 1; j >= i; j-- {
			parent := (j - 1) / 2
			if nodes[j].Priority < nodes[parent].Priority {
				nodes[j], nodes[parent] = nodes[parent], nodes[j]
			}
		}
		nodes[len(nodes)-1], nodes[i] = nodes[i], nodes[len(nodes)-1]
	}
}

func buildMinHeap(nodes []*Node, begin int) {
	for j := len(nodes) - 1; j >= begin; j-- {
		parent := (j - 1) / 2
		if nodes[j].Priority < nodes[parent].Priority {
			nodes[j], nodes[parent] = nodes[parent], nodes[j]
		}
	}
}
