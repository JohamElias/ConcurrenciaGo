package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	ch0 chan struct{} = make(chan struct{})
	ch1 chan struct{} = make(chan struct{})
	ch2 chan struct{} = make(chan struct{})
	ch3 chan struct{} = make(chan struct{})
	ch4 chan struct{} = make(chan struct{})
	ch5 chan struct{} = make(chan struct{})
	ch6 chan struct{} = make(chan struct{})
	mux bool          = false
)

func worker1() {
	for {
		wg.Add(1)
		<-ch1
		fmt.Printf("A")
		wg.Done()
		ch3 <- struct{}{}
	}
}

func worker2() {
	for {
		wg.Add(1)
		<-ch1
		fmt.Printf("B")
		wg.Done()
		ch3 <- struct{}{}
	}
}

func worker3() {
	for {
		wg.Add(1)
		<-ch3
		fmt.Printf("C")
		wg.Done()
		ch4 <- struct{}{}
	}
}

func worker4() {
	for {
		wg.Add(1)
		<-ch4
		fmt.Printf("D")
		wg.Done()
		ch5 <- struct{}{}
	}
}

func worker5() {
	for {
		wg.Add(1)
		<-ch5
		fmt.Printf("E")
		wg.Done()
		ch1 <- struct{}{}
		/*if mux == true {
			mux = false
			ch1 <- struct{}{}
		} else {
			mux = true
			ch2 <- struct{}{}
		}*/
	}
}

func main() {

	go worker1()
	go worker2()
	go worker3()
	go worker4()
	go worker5()
	ch1 <- struct{}{}
	wg.Wait()
}
