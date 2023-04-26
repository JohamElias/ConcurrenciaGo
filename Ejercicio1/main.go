package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	ch1 chan struct{} = make(chan struct{}, 1)
	ch2 chan struct{} = make(chan struct{}, 1)
	ch3 chan struct{} = make(chan struct{}, 1)
)

func worker1() {
	<-ch1
	fmt.Printf("Sistemas ")
	wg.Done()
	ch2 <- struct{}{}
}

func worker2() {
	defer wg.Done()
	fmt.Printf("INF239 ")
	ch1 <- struct{}{}
}

func worker3() {
	defer wg.Done()
	<-ch2
	fmt.Printf("Operativos ")
	ch3 <- struct{}{}
}

func worker4() {
	<-ch3
	defer wg.Done()
	fmt.Printf("\n")
}

func main() {
	wg.Add(4)
	go worker1()
	go worker2()
	go worker3()
	go worker4()
	wg.Wait()
}
