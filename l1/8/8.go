package main 

import (
    "fmt"
)

func main() {
    i := 123456   // 0000 0000 0000 0001 1110 0010 0100 0000
    fmt.Printf("Before change:              %b\n",i)

    // установим 24ый бит в 1
    i |= 1<<23    // 0000 0000 1000 0001 1110 0010 0100 0000
    fmt.Printf("After first change:  %b\n",i) 

    // установим 16ый бит в 0
    i &= ^(1<<15) // 0000 0000 1000 0001 0110 0010 0100 0000
    fmt.Printf("After second change: %b\n",i) 
}
