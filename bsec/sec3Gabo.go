package main

import (
	"fmt"
	"sync"
)

var (
	w     sync.WaitGroup
	count = 0
	ch1   = make(chan int)
	ch2   = make(chan int)
	ch3   = make(chan int)
	ch4   = make(chan int)
	ch5   = make(chan int)
	ch6   = make(chan int)
	ch7   = make(chan int)
	ch8   = make(chan int)
)

func routineA() {
	for {
		//Se la agrega 1 a la espera
		w.Add(1)
		//Espera un mensaje de ch2
		<-ch8

		fmt.Printf("A")
		//Se avisa que terminó el hilo
		w.Done()
		//Pasa un mensaje por ch1
		ch1 <- 1
	}
}

//Lo mismo de la rutina A para el resto de rutinas
func routineB() {
	for {
		w.Add(1)
		<-ch8
		fmt.Printf("B")
		w.Done()
		ch1 <- 1
	}
}

func routineC() {
	for {
		w.Add(1)
		<-ch1
		fmt.Printf("C")
		w.Done()
		ch6 <- 1
	}
}

func routineD() {
	for {
		w.Add(1)
		<-ch6
		fmt.Printf("D")
		w.Done()
		ch7 <- 1
	}
}

func routineE() {
	for {
		w.Add(1)
		<-ch7
		fmt.Printf("E")
		w.Done()
		ch8 <- 1
		/*
			select {
			case <-ch8:
				ch2 <- 1
			case <-ch8:
				ch2 <- 1
			}
		*/
	}
}

/*func manageSelects(){
	select{
		case <-ch3:
			ch4<-1
		case <-ch1:
			ch2<-1
		default:
			manageSelects()
	}
	ch5<-1
	<-ch8
	count ++
	if count == 50 {
		os.Exit(1)
	}
}*/
/*
func manageSignals(){
	for{
		select{
		case <-ch3:
			ch4<-1
			ch5<-1
			<-ch8
		case <-ch1:
			ch2<-1
			ch5<-1
			<-ch8
	}

	count ++
		if count == 50 {
			os.Exit(1)
		}
	}
}*/

func main() {

	go routineA()
	go routineB()
	go routineC()
	go routineD()
	go routineE()
	//Se manda el primer mensaje para accionar una rutina
	//Se escoge en ch2 para empezar a activar A
	//fmt.Printf("EAAAAA")
	ch8 <- 1
	//manageSignals()
	//Se le pide esperar a que acaben todos los hilos
	//Los hilos no acabaran porque están en bulce infinito
	//y por cada bucle, se aumentará una unidad en espera
	//Con x.Add(1)
	//Y si se imprime, se reducirá una unidad de espera
	//Con w.Done()
	w.Wait()
	//fmt.Printf("\n")
}
