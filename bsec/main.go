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
	ch4 chan struct{} = make(chan struct{}, 1)
	ch5 chan struct{} = make(chan struct{}, 1)
)

func worker1() {
	<-ch1
	fmt.Printf("A")
	wg.Done()
	ch2 <- struct{}{}
}

func worker2() {
	<-ch2
	fmt.Printf("B")
	wg.Done()
	ch3 <- struct{}{}
}

func worker3() {
	<-ch3
	fmt.Printf("C")
	wg.Done()
	ch4 <- struct{}{}
}

func worker4() {
	<-ch4
	fmt.Printf("D")
	wg.Done()
	ch5 <- struct{}{}
}

func worker5() {
	<-ch5
	fmt.Printf("E")
	wg.Done()
	ch1 <- struct{}{}
}

func main() {
	ch1 <- struct{}{}
	for true {
		wg.Add(5)
		worker1()
		worker3()
		worker4()
		worker5()
		worker2()
		wg.Wait()
	}

}
