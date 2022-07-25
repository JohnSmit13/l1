package main

import (
	"fmt"
)

func main() {
    str1:= "asdf" // true
    str2:= "asdfF" // false
    str3:= "asdfeNgK" // false
    str4:= "aAdf" // false

    fmt.Printf("str1: %s, res: %v\n",str1, check(str1))
    fmt.Printf("str2: %s, res: %v\n",str2, check(str2))
    fmt.Printf("str3: %s, res: %v\n",str3, check(str3))
    fmt.Printf("str4: %s, res: %v\n",str4, check(str4))
}

func check(s string) bool {

    m := make(map[byte]int)
    c := make([]byte,len(s))

    // преводим каждый символ к нижнему регистру
    for i := range s {
        if s[i] >= 0x61 && s[i] <= 0x7A {
            c[i] = s[i]
        } else {
            // разница в значениях между верхним и нижним регистром = 0x20(32)
            c[i] = s[i] + 0x20
        }

        if m[c[i]] != 0 {
            return false
        } else {
            m[c[i]]++
        }
    }
    return true
}
