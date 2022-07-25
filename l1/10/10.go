package main 

import (
    "fmt"
)

func main() {
    arr := []float32{-25.4, 4.4, 5.6, 6.2, 1.0, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
    res := make(map[int][]float32)

    for _,v := range arr {
        var tmp int // ключ
        var tmpV float32
        if v < 0 {
            tmpV = v - 10  // значения от -30 до -20 будут иметь ключ -30, от -20 до -10 : -20 и т.д.
        } else { tmpV = v }// 
        tmp = int(tmpV/10) // при округлении отбрасывается дробная часть: -25.4/10 = -2!
        tmp *= 10          // возвращаем потерянную в ходе округления десятку
        res[tmp] = append(res[tmp],v)
    }
    for i,v :=range res {
        fmt.Printf("Key: %3d, value: %v\n",i,v)
    }
}
