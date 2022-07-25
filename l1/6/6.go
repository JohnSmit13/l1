package main

import (
    "fmt"
    "time"
    "sync"
)

// способ 1: просто подождать пока горутина сама завершится
//           если она не в безконечном цикле.

// способ 2: прекратить работу горутины завершив работу главной горутины.
// способ 3: послать сигнал горутине, но для этого она должна проверять канал 
//           по которому мы собираемся послать этот сигнал.

func gorutine1(wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("done")
    time.Sleep(time.Second * 1) // завершится сама
}

func gorutine2() {
    defer fmt.Println("done") // не увидим
    for {} // никогда сама не остановится.
}

func gorutine3(c chan int) {
    defer fmt.Println("done")
    for {
        time.Sleep(time.Millisecond * 500)
        select {
            case <-c:
                return
        }
    }
}
//1
/*func main() {
    var wg sync.WaitGroup
    fmt.Printf("start gorutine1\n")
    wg.Add(1)
    go gorutine1(&wg)
    wg.Wait()
}*/

//2
/*func main() {
    go gorutine2()
    time.Sleep(time.Second *1)
}*/

//3
func main() {
    c := make(chan int)
    go gorutine3(c)
    c <- 1
    time.Sleep(time.Second * 1)
}
