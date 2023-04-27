package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	ch0 chan struct{} = make(chan struct{}, 1)
	ch1 chan struct{} = make(chan struct{}, 1)
	ch2 chan struct{} = make(chan struct{}, 1)
	ch3 chan struct{} = make(chan struct{}, 1)
	ch4 chan struct{} = make(chan struct{}, 1)
	ch5 chan struct{} = make(chan struct{}, 1)
)

func worker1() {
	<-ch0
	//ch3 <- struct{}{}
	<-ch1
	fmt.Printf("A")
	ch3 <- struct{}{}
	ch2 <- struct{}{}
	wg.Done()
}

func worker2() {
	<-ch0
	//ch3 <- struct{}{}
	<-ch2
	fmt.Printf("B")
	ch3 <- struct{}{}
	ch1 <- struct{}{}
	wg.Done()
}

func worker3() {
	<-ch3
	fmt.Printf("C")
	ch4 <- struct{}{}
	wg.Done()
}

func worker4() {
	<-ch4
	fmt.Printf("D")
	ch5 <- struct{}{}
	wg.Done()
}

func worker5() {
	<-ch5
	fmt.Printf("E")
	ch0 <- struct{}{}
	wg.Done()
}

func main() {
	ch0 <- struct{}{}
	ch1 <- struct{}{}
	for true {
		wg.Add(5)
		go worker1()
		go worker2()
		go worker3()
		go worker4()
		go worker5()
		wg.Wait()
	}
}
