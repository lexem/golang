package main

const p = 1000000007
const xConst = 257

type Stringer struct {
	x []int
	h []int
}

func NewStringer(str string) *Stringer {
	runes := []rune(" " + str)
	n := len(str)

	x := make([]int, n+1)
	h := make([]int, n+1)
	x[0] = 1

	for i := 1; i < n+1; i++ {
		h[i] = (h[i-1]*xConst + int(runes[i])) % p
		x[i] = (x[i-1] * xConst) % p
	}

	return &Stringer{
		x: x,
		h: h,
	}
}

func (s *Stringer) IsEqual(a int, b int, len int) bool {
	return (s.h[a+len]+s.h[b]*s.x[len])%p ==
		(s.h[b+len]+s.h[a]*s.x[len])%p
}
