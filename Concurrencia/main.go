package main

import (
	"fmt"
	"sync"
)

func worker() {
	for x := 0; x < 1000000; x++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
	wg.Done()
}

var (
	count int
	wg    sync.WaitGroup
	mu    sync.Mutex
)

func main() {
	for x := 0; x < 5; x++ {
		wg.Add(1)
		go worker()
	}
	wg.Wait()
	fmt.Printf("El valor esperado de count es: 5000000 y el valor final es %d\n", count)
}
