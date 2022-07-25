package main 

import (
    "fmt"
)

func main() {
    x,y := 100, 200
    fmt.Printf("x=%d, y=%d\n",x,y)

    // go style
    y,x = x,y
    fmt.Printf("x=%d, y=%d\n",x,y)

    // binary magic)
    x ^= y
    y = x^y
    x = x^y
    fmt.Printf("x=%d, y=%d\n",x,y)
}
