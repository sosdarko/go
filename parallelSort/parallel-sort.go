package main

import (
	"fmt"
)

// merges two sorted sequences into new sorted one
func merge(a []int, b []int) (c []int) {
	i := 0
	j := 0
	// will be placed on the heap
	c = make([]int, len(a)+len(b))
	// run until both sequences are read
	for i < len(a) || j < len(b) {
		if i >= len(a) {
			c[i+j] = b[j]
			j++
		} else {
			if j >= len(b) {
				c[i+j] = a[i]
				i++
			} else {
				if a[i] < b[j] {
					c[i+j] = a[i]
					i++
				} else {
					c[i+j] = b[j]
					j++
				}
			}
		}

	}
	fmt.Printf("i=%d, j=%d", i, j)
	return
}

func main() {
	a := []int{1, 2, 4, 10}
	b := []int{3, 5, 7}
	c := merge(a, b)
	fmt.Println(c)
}
