package main

import (
    "fmt"
    "sync"
)

func routine(n int){
    defer wg.Done()
    fmt.Printf("I'm goroutine %d\n",n)
}
var wg sync.WaitGroup
func main(){
    for x:=0;x<5;x++{
        wg.Add(1)
        go routine(x)
    }
    wg.Wait()
}

