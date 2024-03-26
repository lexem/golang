package main

type Village struct {
	index int
	time  int
}

type VillageHeap []Village

func (h VillageHeap) Len() int           { return len(h) }
func (h VillageHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h VillageHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *VillageHeap) Push(x any) { *h = append(*h, x.(Village)) }

func (h *VillageHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
