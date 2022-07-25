package main

import (
    "fmt"
)

func main() {
    x := []int{1,2,3,4,5,6,7,8,9,10,11,12}
    fmt.Println(x)
    i := 5
    for j := 0; j < len(x) - 1 - i; j++ {
        x[i+j] = x[i+j+1]
    }
    x = x[:len(x)-1]
    fmt.Println(x)
}
