package main 

import (
    "fmt"
)

func main() {
    arr := []float32{-25.4, 4.2, 4.4, 5.6, 6.2, 1.0, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
    arr2 := []float32{-25.4, 13.0, 15.5, -21.0, 4.2, 10.5, 11.9, 42.9}

    res := make(map[float32]int)

    for _,v := range arr {
        res[v]++
    }

    for _,v := range arr2 {
        res[v]++
    }

    var cross []float32
    for i,v :=range res {
        if v > 1{
            cross = append(cross,i)
        }
    }

    fmt.Printf("Cross:\n%v\n",cross)
}

// вариант 2
/*func main() {
    arr := []float32{-25.4, 4.2, 4.4, 5.6, 6.2, 1.0, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
    arr2 := []float32{-25.4, 13.0, 15.5, -21.0, 4.2, 10.5, 11.9, 42.9}
    var cross []float32

    for _,v1 := range arr {
        for _,v2 := range arr2 {
            if v1 == v2 {
                cross = append(cross,v1)
            }
        }
    }
    fmt.Printf("Cross:\n%v\n",cross)
}*/
