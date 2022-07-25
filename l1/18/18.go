package main

import (
	"fmt"
    "sync"
    "time"
)

var stime time.Duration = time.Millisecond * 100
var myCounter counter

type counter struct {
    value uint64
    mx  sync.Mutex
    mx2 chan int // вместо блокировки мьютекса мы можем использовать канал с размером буфера 1
}

func main() {
    myCounter.mx2 = make(chan int,1)
    go increaser(&myCounter)
    go increaser(&myCounter)
    time.Sleep(time.Millisecond * 1000)
    fmt.Printf("counter now is: %d\n",myCounter.value)
}

func(c *counter) Inc() {
    fmt.Println("Loking mutex")
    c.mx.Lock()
    fmt.Println("Increasing value")
    c.value++
    fmt.Println("Unloking mutex")
    fmt.Println("   ")
    c.mx.Unlock()
}

func(c *counter) Inc2() {
    fmt.Println("Loking \"mutex\"")
    c.mx2 <- 1
    fmt.Println("Increasing value")
    c.value++
    fmt.Println("Unloking \"mutex\"")
    fmt.Println("   ")
    <- c.mx2
}

func increaser(x *counter) {
    for i := 0; i < 10; i++ {
        //x.Inc()
        x.Inc2()
        time.Sleep(stime)
    }
}
