package main

import (
	"fmt"
	"sync"
)

var (
	buffer [5]int = [5]int{-1, -1, -1, -1, -1}
	index  int
	wg     sync.WaitGroup
	mu     sync.Mutex
)

func productor(c chan int) {
	for n := 0; n < 20; n++ {
		item := n * n
		index = n % 5
		mu.Lock()
		buffer[index] = item
		mu.Unlock()
		fmt.Printf("productor %d %d %v\n", index, item, buffer)
		c <- index
	}
	//close(c)
	wg.Done()
}

func consumidor(c chan int) {
	var item int
	for index := range c {
		mu.Lock()
		item = buffer[index]
		buffer[index] = -1
		mu.Unlock()
		fmt.Printf("Consumidor %d %d %v\n", index, item, buffer)
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	wg.Add(2)
	go consumidor(c)
	go productor(c)
	wg.Wait()
}
