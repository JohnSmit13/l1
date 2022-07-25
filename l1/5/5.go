package main

import (
    "fmt"
    "time"
)

func main() {

    ch  := make(chan int)

    go func(i chan int) {
        for {
            i <- 4 // why not?))
            time.Sleep(time.Millisecond * 100)
        }
    }(ch)

    go func(o chan int) {
        for {
            fmt.Println("got int:",<-o)
        }
    }(ch)

    time.Sleep(time.Second * 5)
}
