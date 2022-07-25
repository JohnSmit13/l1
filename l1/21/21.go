package main

import (
    "fmt"
)

type drawer interface {
    Draw()
}

type badDrawer struct {}

type adaptedDrawer struct {
    badDrawer
    drawer
}

func main() {

    var i drawer
    var d badDrawer
    d.BadDraw()
    fmt.Println()

    //i = d  // так нельзя, d не соответствует интерфейсу
    var a adaptedDrawer
    i = a
    i.Draw()
}

func(d badDrawer) BadDraw() {
    fmt.Println("BadDraw")
}

func(d adaptedDrawer) Draw() {
    d.BadDraw()
    fmt.Println("GoodDraw!")
}
