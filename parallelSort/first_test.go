package main

import (
	"fmt"
	"testing"
)

func TestTimeConsuming(t *testing.T) {
	DoIt(15000000, 150000)
}

func TestInPlaceMerge(t *testing.T) {
	a := [11]int{1, 3, 5, 7, 10}
	b := [6]int{-3, -1, 6, 6, 13, 25}
	inPlaceMerge(a[:5], b[:])
	fmt.Println(a)
}

func TestBigInPlaceMerge(t *testing.T) {
	s := 150000000
	c := 1500000
	fmt.Printf("sorting %d integers with chunk size %d\n", s, c)
	DoIt2(s, c, false)
}
