package main

import (
	"container/heap"
	"errors"
	"fmt"
	"log"
)

type PresentHeap struct {
	Value int
	Size  int
}

type IntHeap []PresentHeap

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	if h[i].Value == h[j].Value {
		return h[i].Size < h[j].Size
	}
	return h[i].Value > h[j].Value
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(PresentHeap))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getNCoolestPresents(a []PresentHeap, n int) ([]interface{}, error) {
	h := IntHeap(a)
	var pr []interface{}
	if n < 0 {
		return nil, errors.New("n cannot be negative")
	}
	if n > len(a) {
		return nil, errors.New("n more gifts")
	}
	heap.Init(&h)
	for i := 0; i < n; i++ {
		pr = append(pr, heap.Pop(&h))
	}
	return pr, nil
}

func main() {

	tmp := []PresentHeap{
		PresentHeap{Value: 5, Size: 1},
		PresentHeap{Value: 4, Size: 5},
		PresentHeap{Value: 3, Size: 1},
		PresentHeap{Value: 5, Size: 2},
	}
	res, err := getNCoolestPresents(tmp, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
