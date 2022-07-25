package main 

import (
    "fmt"
    "time"
)

func main() {
    numbers := make(chan uint64)
    squares := make(chan uint64)

    go producer(numbers)
    go calculator(numbers,squares)
    go printer(squares)

    for {}
}

// генерирует числа
func producer(ch chan uint64) {
    var i uint64
    for {
        ch <- i
        i++
        time.Sleep(time.Millisecond * 250)
    }
}

// считает крадраты чисел
func calculator(in,out chan uint64) {
    for {
        x := <- in
        out <- x*x
    }
}

// выводит квадраты на экран
func printer(ch chan uint64) {
    for {
        fmt.Printf("Square: %d\n", <- ch)
    }
}
