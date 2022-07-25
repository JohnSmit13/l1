package main

import (
	"fmt"
)

func main() {

    var d interface{} = 34
    wichType(d)
    d = "value"
    wichType(d)
    d = true
    wichType(d)
    d = make(chan interface{})
    wichType(d)
    d = make(chan int)
    wichType(d)
}


func wichType(x interface{}) {
    switch x.(type) {
        case string:
            fmt.Println("String")
        case int:
            fmt.Println("Int")
        case bool:
            fmt.Println("Bool")
        case chan interface{}:
            fmt.Println("Chanel")
        default: 
            fmt.Println("something else")
    }
}
