package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	buffer   [5]int = [5]int{-1, -1, -1, -1, -1}
	index    int
	wg       sync.WaitGroup
	mu       sync.Mutex
	c        chan int = make(chan int, 5)
	ch1, ch2 chan int = make(chan int), make(chan int)
)

func productor() {
	for n := 0; n < 20; n++ {
		item := n * n
		//mu.Lock()
		index = n % 5
		if buffer[index] != -1 {
			ch1 <- 1
		}
		c <- item
		mu.Lock()
		buffer[index] = item
		mu.Unlock()
		//mu.Unlock()
		fmt.Printf("productor %d %d %v\n", index, item, buffer)
	}
	wg.Done()
}

func consumidor() {
	var item int
	for n := 0; n < 20; n++ {
		<-c
		//mu.Lock()
		index = n % 5
		for buffer[index] == -1 {
			index++
			if index == 5 {
				index = 0
			}
		}
		item = buffer[index]
		mu.Lock()
		buffer[index] = -1
		mu.Unlock()
		select {
		case <-ch1:

		default:
			time.Sleep(time.Millisecond)
		}
		//mu.Unlock()
		fmt.Printf("Consumidor %d %d %v\n", index, item, buffer)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go consumidor()
	go productor()
	wg.Wait()
}
