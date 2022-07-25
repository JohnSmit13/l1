package main

import (
	"fmt"
	"sync"
)

func main() {

    const arrSize = 5

    arr := [arrSize]int{2,4,6,8,10}
    var wg sync.WaitGroup
    c := make(chan int,arrSize)

    for _,v := range arr {
        go square(v,&wg,c)
        wg.Add(1)
    }
    wg.Wait()
    close(c)
    var res int
    for i := range c {
        res += i
    }
    fmt.Println("result:",res) //220
}

func square(v int, wg *sync.WaitGroup, c chan int) {
    c <- v*v
    wg.Done()
}
