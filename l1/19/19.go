package main

import (
	"fmt"
)

func main() {
    str:= "главрыба"
    fmt.Printf("string:         %s\n",str)

    r:= []rune(str) // "превращаем" строку в слайс рун
    end:=len(r)-1

    // каждая руна занимет фиксированное количество байт(4),
    // поэтому мы можем поменять любые 2 руны местами не повредив данные
    for i := 0; i <= end/2; i++ {
        r[i],r[end-i] = r[end-i],r[i]
    }

    str = string(r) // заменяем исходную строку
    fmt.Printf("reverse string: %s\n",str)
}
