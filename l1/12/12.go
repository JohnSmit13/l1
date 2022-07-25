package main 

import (
    "fmt"
)

func main() {
    arr := []string{"cat", "cat", "dog", "cat", "tree"}
    res := make(map[string]int)

    for _,v := range arr {
        res[v]++
    }

    var ans []string
    for i := range res {
        ans = append(ans,i)
    }

    fmt.Printf("Answer:\n%v\n",ans)
}
