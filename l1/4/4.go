package main

import (
    "fmt"
    "sync"
    "os"
    "time"
    "os/signal"
)

func main() {

    var N,i uint

    fmt.Println("Enter the number of workers")
    fmt.Fscan(os.Stdin, &N)

    ch := make(chan int, N)
    var wg sync.WaitGroup

    wg.Add(int(N)) // I have question about this
    for i = 1; i <= N; i++ {
        go worker(i,ch,&wg)
        //wg.Add(1) // some times this doesn't work correctly...???
    }

    osSig := make(chan os.Signal, 1)
    signal.Notify(osSig, os.Interrupt)
    j := 0
    for {
        if len(osSig) < 1 {

            fmt.Println("sending integer:",j)
            ch <- j; j++
            time.Sleep(time.Millisecond * 100)

        } else {

            close(ch)
            break;
        }
    }
    wg.Wait()
    fmt.Println("All workers correctly done their work!")
}

func worker(id uint, c chan int, wg *sync.WaitGroup) {
    defer fmt.Printf("Worker %d done his work!\n", id)
    for {
        if i,ok := <- c; ok {
            fmt.Printf("Worker %d got integer: %d\n",id, i)
        } else {
            break;
        }
    }
    wg.Done()
}
