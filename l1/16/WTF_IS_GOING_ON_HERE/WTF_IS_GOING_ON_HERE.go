package main

import (
	"fmt"
    "time"
)

func main() {
    fmt.Printf("hi\n")
    a := []int{9,3,45,2,6,3,43,1,6,24,53,14,23,54,34,62,43,55,23,53,12,95}
    fmt.Printf("%v\n",quickSort(a))
}

func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
     
    beg, end := 0, len(arr)-1
    var sleepTime time.Duration = 100

    //for range arr {
    for beg < end {
        fmt.Printf("  \n")
        fmt.Printf("beg: %d, end: %d\n",beg,end)
        fmt.Printf("%t\n",beg<end)
        if arr[beg] > arr[end] {
            fmt.Printf("arr[beg]: %d, arr[end]: %d\n",arr[beg],arr[end])
            arr[beg], arr[end] = arr[end], arr[beg]
            fmt.Printf("arr[beg]: %d, arr[end]: %d\n",arr[beg],arr[end])
            beg++
            fmt.Printf("arr[beg]: %d, arr[end]: %d\n",arr[beg],arr[end])
            time.Sleep(time.Millisecond * sleepTime)
            fmt.Printf("change!    beg: %d, end: %d\n",beg,end)
        } else {
            fmt.Printf("arr[beg]: %d, arr[end]: %d\n",arr[beg],arr[end])
            end--
            fmt.Printf("arr[beg]: %d, arr[end]: %d\n",arr[beg],arr[end])
            time.Sleep(time.Millisecond * sleepTime)
            fmt.Printf("no change! beg: %d, end: %d\n",beg,end)
            if beg == end {fmt.Printf("STOP!!!");break}
        }
    }
    quickSort(arr[:beg])
    quickSort(arr[beg:])
    return arr
}
