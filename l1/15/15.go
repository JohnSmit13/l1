package main 

var justString string
func someFunc() {
    v := createHugeString(1 << 10)
    // чтобы избежать утечки памяти нужно создать новый слайс 
    // и явно скопировать все элементы
    v2:= make([]byte,100)
    for i := range v2 {
        v2[i] = v[i]
    }
    justString = string(v2)
    // память занятая v освободится
}

func main() {
    someFunc()
}

func createHugeString(i int) string {
    res := make([]byte,i)
    return string(res)
}
