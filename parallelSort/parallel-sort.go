package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// inPlaceMerge supposes that aTo has enough capacity to take aFrom array in
func inPlaceMerge(aTo []int, aFrom []int) {
	i := len(aTo) - 1
	j := len(aFrom) - 1
	for j >= 0 {
		startI := i
		// find place from j-th element of aFrom in aTo
		for aFrom[j] < aTo[i] {
			i--
			if i < 0 {
				break
			}
		}
		// shift members of aTo for j position right
		copy(aTo[i+j+2:startI+j+2], aTo[i+1:startI+1])
		if i >= 0 {
			// this could be further optimized
			// by traversing aFrom (reducing j) to find least memeber that is bigger than aTo[i]
			// and then copy whole sequence at once
			copy(aTo[i+j+1:i+j+2], aFrom[j:j+1])
		} else {
			copy(aTo[:j+1], aFrom[:j+1])
			break
		}
		j--
	}
}

// Merge merges two sorted sequences into new sorted one
// note that result slice will be put on heap as out param
func Merge(a []int, b []int) (c []int) {
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
		//fmt.Printf("chunk received at %v\n", time.Now())
		if i == 1 {
			sortedArray = sortedSubarray
		} else {
			// merge so far sorted array with new chunk
			sortedArray = Merge(sortedArray, sortedSubarray)
		}
	}
	endTime := time.Now()
	fmt.Printf("all done at %v\n", endTime)
	fmt.Printf("execution time = %v\n", time.Duration(endTime.UnixNano()-startTime.UnixNano()))
}

// DoIt2 ...
func DoIt2(nSize, nChunk int, bPrint bool) {
	ar := createRandomArray(nSize)
	arrayChannel := make(chan []int)
	if bPrint {
		fmt.Println("Original array:")
		fmt.Println(ar)
	}
	sortedArray := make([]int, nSize)
	startTime := time.Now()
	fmt.Printf("starting go routines at %v\n", startTime)
	for i := 0; i < nSize; i += nChunk {
		// in place sort of subarray
		go smallSort(ar[i:i+nChunk], arrayChannel)
	}
	fmt.Printf("receiving go routines at %v\n", time.Now())
	for i := 0; i < nSize; i += nChunk {
		sortedSubarray := <-arrayChannel
		//fmt.Printf("chunk received at %v\n", time.Now())
		if i == 0 {
			copy(sortedArray, sortedSubarray)
		} else {
			// merge so far sorted array with new chunk
			inPlaceMerge(sortedArray[:i], sortedSubarray)
		}
	}
	endTime := time.Now()
	if bPrint {
		fmt.Println(sortedArray)
	}
	fmt.Printf("all done at %v\n", endTime)
	fmt.Printf("execution time = %v\n", time.Duration(endTime.UnixNano()-startTime.UnixNano()))
}

func main() {
	DoIt2(100, 10, true)
}
