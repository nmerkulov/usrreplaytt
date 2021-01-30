package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	intCount       = 20
	minInt         = 1
	maxInt         = 10000
	consumersCount = 10
)

func main() {
	rand.Seed(time.Now().Unix())
	ch := generateInts(intCount, func() int {
		return rand.Intn(maxInt-minInt) + minInt
	})
	sem := make(chan struct{}, consumersCount)
	var wg sync.WaitGroup
	for i := range ch {
		sem <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
			fmt.Printf("sleep for a %d milliseconds\n", i)
			time.Sleep(time.Duration(i) * time.Millisecond)
		}(i)
	}
	wg.Wait()
}

func generateInts(count int, f func() int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- f()
		}
	}()
	return ch
}
