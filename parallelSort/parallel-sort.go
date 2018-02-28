package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
func inPlaceMerge(mergeTo []int, mergeFrom []int) {
	i := 0
	j := 0
}
*/
// merges two sorted sequences into new sorted one
func merge(a []int, b []int) (c []int) {
	i := 0
	j := 0
	// will be placed on the heap
	c = make([]int, len(a)+len(b))
	// run until both sequences are read
	for i < len(a) || j < len(b) {
		// if a is read, fill rest with b
		if i >= len(a) {
			c[i+j] = b[j]
			j++
		} else {
			// if b is read, fill rest with a
			if j >= len(b) {
				c[i+j] = a[i]
				i++
			} else {
				// if both are not read, compare and move one that has smaller value
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
	return
}

func smallSort(a []int, ss chan []int) {
	// here we use go's sort package
	sort.Ints(a)
	ss <- a
}

func createRandomArray(nSize int) (b []int) {
	b = make([]int, nSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nSize; i++ {
		b[i] = rand.Intn(10000) - 500
	}
	// returns named return parameter b
	return
}

// DoIt will do it
func DoIt(nSize, nChunk int) {
	ar := createRandomArray(nSize)
	arrayChannel := make(chan []int)
	//fmt.Println("Original array:")
	// fmt.Println(ar)
	var sortedArray []int
	startTime := time.Now()
	fmt.Printf("starting go routines at %v\n", startTime)
	for i := 0; i < nSize; i += nChunk {
		// in place sort of subarray
		go smallSort(ar[i:i+nChunk], arrayChannel)
	}
	fmt.Printf("receiving go routines at %v\n", time.Now())
	for i := 0; i < nSize; i += nChunk {
		sortedSubarray := <-arrayChannel
		fmt.Printf("chunk received at %v\n", time.Now())
		if i == 1 {
			sortedArray = sortedSubarray
		} else {
			// merge so far sorted array with new chunk
			sortedArray = merge(sortedArray, sortedSubarray)
		}
	}
	endTime := time.Now()
	fmt.Printf("all done at %v\n", endTime)
	fmt.Printf("execution time = %v\n", time.Duration(endTime.UnixNano()-startTime.UnixNano()))
}

func main() {
	DoIt(15000000, 15000)
}
