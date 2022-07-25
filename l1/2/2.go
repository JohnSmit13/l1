package main

import (
	"fmt"
	"sync"
)

func main() {

    arr := [5]int{2,4,6,8,10}
    var wg sync.WaitGroup

    for _,v := range arr {
        go square(v,&wg)
        wg.Add(1)
    }
    wg.Wait()
}
func square(v int, wg *sync.WaitGroup) {
    fmt.Println(v*v)
    wg.Done()
}

//вариант с использованием канала
/*func main() {

    arr:= [5]int{2,4,6,8,10}
    c:= make(chan int,5)
    var wg sync.WaitGroup

    for _,v := range arr {
        wg.Add(1)
        go square(c,&wg,v)
    }
    wg.Wait()
    close(c)

    for i := 0; i<cap(c); i++ {
        fmt.Println(<-c)
    }
}

func square(c chan int, wg *sync.WaitGroup, v int) {
    c <- v*v
    wg.Done()
}*/
