package main

import (
    "fmt"
    "sync"
    "time"
    "math/rand"
)

var mx sync.Mutex // for guard map

func main() {

    m := make(map[int]string)

    go addValue("first" , &m)
    go addValue("second", &m)

    time.Sleep(time.Second * 2)

    mx.Lock()
    for i,v := range m {
        fmt.Printf("key: %d, value: %s\n", i,v)
    }
}

func addValue(n string, m *map[int]string) {
    i := 0
    for {
        mx.Lock()
        fmt.Println("Mutex is locked!")
        (*m)[i] = n
        fmt.Printf("%s value: %d\n",n,i)
        mx.Unlock()
        fmt.Println("Mutex is unlocked!")
        fmt.Println("  ")
        i++
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
    }
}
