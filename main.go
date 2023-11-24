package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Batch struct {
	Begin int
	End   int
}

func main() {

	//nArr := []int{10, 100, 1000, 100000, 1000000, 10000000, 100000000, 1000000000}
	nArr := []int{10, 100, 1000, 100000, 1000000}
	mArr := []int{1, 2, 3, 4, 5, 10, 20, 30, 100}

	for _, n := range nArr {
		arr, fileT := generateFile(n)
		for _, m := range mArr {
			t := calcBathes(makeBatches(n, m), arr)
			fmt.Printf("n = %d m = %d t = %d\n", n, m, t+fileT)
		}
		arr = nil
	}
}

func generateFile(n int) ([]int, int64) {
	startTime := time.Now()
	filename := "numbers-" + strconv.FormatInt(int64(n), 10) + ".txt"
	res := make([]int, n)
	file, err := os.Create(filename)

	if err != nil {
		fmt.Printf("Error while creating file")
		return res, time.Since(startTime).Milliseconds()
	}

	numbers := ""

	for i := 1; i <= n; i++ {
		if i == 1 {
			numbers += strconv.FormatInt(int64(i), 10)
		} else {
			numbers += " " + strconv.FormatInt(int64(i), 10)
		}
	}

	stringLen, err := file.WriteString(numbers)

	if err != nil {
		fmt.Printf("Error while writing file")
		return res, time.Since(startTime).Milliseconds()
	}

	openFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error while reading file")
		return res, time.Since(startTime).Milliseconds()
	}

	data := make([]byte, stringLen)

	for {
		_, err := openFile.Read(data)
		if err != io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error while reading file")
			break
		}
	}

	strArr := strings.Fields(string(data))

	for i, number := range strArr {
		intNumber, err := strconv.Atoi(number)

		if err != nil {
			fmt.Printf("Error while convert number")
			break
		}

		res[i] = intNumber
	}

	return res, time.Since(startTime).Milliseconds()
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

func calcBatch(batch Batch, arr []int) {
	batchArr := make([]float64, batch.End-batch.Begin+1)

	for j, i := 0, batch.Begin; i <= batch.End; i, j = i+1, j+1 {
		batchArr[j] = float64(arr[i-1])
	}

	for i := 0; i < len(batchArr); i++ {
		batchArr[i] *= math.Pi
	}

	batchArr = nil
}

func calcBathes(batches []Batch, arr []int) int64 {
	var wg sync.WaitGroup
	startTime := time.Now()
	for _, batch := range batches {
		wg.Add(1)
		go func(batch2 Batch) {
			calcBatch(batch2, arr)
			wg.Done()
		}(batch)
	}
	wg.Wait()
	return time.Since(startTime).Milliseconds()
}
