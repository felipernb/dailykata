package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	value uint8
	priority int
	index int
	leftChild *Node
	rightChild *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	a := *pq
	n := len(a)
	a = a[0 : n+1]
	item := x.(*Node)
	item.index = n
	a[n] = item
	*pq = a
}

func (pq *PriorityQueue) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	item.index = -1
	*pq = a[0 : n-1]
	return item
}

func (pq PriorityQueue) Huffman(text string) map[uint8]string {
	symbols := make(map[uint8]int)
	for i := 0; i < len(text); i++ {
		symbols[text[i]]++
	}

	for char, freq := range symbols {
		item := &Node{
			value: char,
			priority: freq,
		}
		heap.Push(&pq, item)
	}

	for len(pq) > 1 {
		rightChild, leftChild := heap.Pop(&pq).(*Node), heap.Pop(&pq).(*Node)
		parent := &Node{
			value : '\b',
			priority: rightChild.priority + leftChild.priority,
			leftChild: leftChild,
			rightChild: rightChild,
		}
		heap.Push(&pq, parent)
	}

	root := heap.Pop(&pq).(*Node)
	table := make(map[uint8]string)
	var symbolsTable func(*Node, string)
	symbolsTable = func(node *Node, prefix string) {
		if node.value != '\b' {
			table[node.value] = prefix
		} else {
			symbolsTable(node.leftChild, prefix + "0")
			symbolsTable(node.rightChild, prefix + "1")
		}
	}
	symbolsTable(root, "")

	return table
}

func GenerateTable(text string) map[uint8]string {
	pq := make(PriorityQueue, 0, len(text)*2)
	return pq.Huffman(text)
}

func main() {
	for k, v := range GenerateTable("this is an example of a huffman tree") {
		fmt.Printf("%c %s\n", k, v)
	}
}
