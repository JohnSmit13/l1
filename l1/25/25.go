package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("going to sleep!")
    sleepPoint := time.Now()
    mySleep(time.Second * 3)
    fmt.Printf("I woke up! I slept %d nanoseconds\n",time.Since(sleepPoint))
}


func mySleep(s time.Duration) {
    wake := time.Duration(time.Now().UnixNano()) + s
    for time.Duration(time.Now().UnixNano()) < wake {
    }
}
