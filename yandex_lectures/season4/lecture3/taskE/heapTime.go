package main

type Time struct {
	index int
	time  float64
}

type TimeHeap []Time

func (h TimeHeap) Len() int           { return len(h) }
func (h TimeHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h TimeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TimeHeap) Push(x any) { *h = append(*h, x.(Time)) }

func (h *TimeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
