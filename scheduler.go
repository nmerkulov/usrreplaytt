package main

import "fmt"

func generateInts(count int, f func() int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- f()
			fmt.Println(i)
		}
	}()
	return ch
}
