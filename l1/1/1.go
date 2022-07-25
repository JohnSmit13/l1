package main

import (
    "fmt"
)

type Human struct {
    name string
}
type Action struct {
    Human //встраивание Human в Action
}


func main() {
    //создаем переменную типа Action
    a := Action{Human{name: "Bob"}}
    //вызываем метод объявленный для Human через переменную типа Action
    a.SayHello()
}

//метод SayHello объявлян для типа Human
func(h Human) SayHello() {
    fmt.Printf("Hello, %s!\n",h.name)
}
