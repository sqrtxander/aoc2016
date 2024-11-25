package utils

import "log"

type MaxHeap []int

func (h MaxHeap) Push(num int) MaxHeap {
	currIdx := len(h)
	parentIdx := (currIdx - 1) / 2
	h = append(h, num)
	for currIdx > 0 && h[currIdx] > h[parentIdx] {
		h[currIdx], h[parentIdx] = h[parentIdx], h[currIdx]
		currIdx = parentIdx
		parentIdx = (currIdx - 1) / 2
	}
	return h
}

func (h MaxHeap) Pop() (MaxHeap, int) {
	if len(h) == 0 {
		log.Fatalln("Attempt to pop from an empty heap")
	}
	result := h[0]
	h[0] = h[len(h)-1]
	h = h[len(h)-1:]

	return MaxHeapify(h, 0), result
}

func (h MaxHeap) Peek() int {
	if len(h) == 0 {
		log.Fatalln("Attempt to peek at an empty heap")
	}
	return h[0]
}

func MaxHeapify(arr []int, idx int) MaxHeap {
	n := len(arr)
	if n == 0 {
		return MaxHeap{}
	}
	hi := idx
	l := 2*idx + 1
	r := 2*idx + 2
	if l < n && arr[hi] < arr[l] {
		hi = l
	}
	if r < n && arr[hi] < arr[r] {
		hi = r
	}
	arr[hi], arr[idx] = arr[idx], arr[hi]
	return MaxHeapify(arr, hi)
}
