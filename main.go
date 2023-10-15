package main

import (
	"fmt"
	"time"
)

func main() {

	nArr := []int{10, 100, 1000, 100000, 1000000, 10000000, 100000000, 1000000000}
	mArr := []int{2, 3, 4, 5, 10, 20, 30, 100}

	for _, n := range nArr {
		arr := make([]int, 0, n)
		for i := 1; i <= n; i++ {
			arr = append(arr, i)
		}
		for _, m := range mArr {
			t := time.Now()

			border := make(chan struct{}, m)
			for i := range arr {
				border <- struct{}{}
				index := i
				go func() {
					// some func with number
					arr[index] *= arr[index]
				}()
				<-border
			}

			fmt.Printf("n=%d m=%d t=%d\n", n, m, time.Since(t).Milliseconds())
		}
	}
}
