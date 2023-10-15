package main

import (
	"fmt"
	"sync"
	"time"
)

type Batch struct {
	Begin int
	End   int
}

func main() {

	nArr := []int{10, 100, 1000, 100000, 1000000, 10000000, 100000000, 1000000000}
	mArr := []int{2, 3, 4, 5, 10, 20, 30, 100}

	for _, n := range nArr {
		for _, m := range mArr {
			t := calcBathes(makeBatches(n, m))
			fmt.Printf("n = %d m = %d t = %d\n", n, m, t)
		}
	}
}

func makeBatches(n, m int) []Batch {
	if n < m {
		m = n
	}
	batches := make([]Batch, m)

	itemsCount := n / m
	div := n % m

	for i := 0; i < m; i++ {
		start := (i * itemsCount) + 1

		if i+1 >= m {
			batches[i] = Batch{
				Begin: start,
				End:   start + itemsCount - 1 + div,
			}
		} else {
			batches[i] = Batch{
				Begin: start,
				End:   start + itemsCount - 1,
			}
		}
	}

	return batches
}

func calcBatch(batch Batch) {
	arr := make([]int, batch.End-batch.Begin+1)

	for j, i := 0, batch.Begin; i <= batch.End; i, j = i+1, j+1 {
		arr[j] = i
	}

	for i := 0; i < len(arr); i++ {
		arr[i] *= arr[i]
	}

	arr = nil
}

func calcBathes(batches []Batch) int64 {
	var wg sync.WaitGroup
	startTime := time.Now()
	for _, batch := range batches {
		wg.Add(1)
		go func(batch2 Batch) {
			calcBatch(batch2)
			wg.Done()
		}(batch)
	}
	wg.Wait()
	return time.Since(startTime).Milliseconds()
}
