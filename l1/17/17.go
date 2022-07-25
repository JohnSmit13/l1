package main

import (
	"fmt"
)

func main() {
    // создаем массив в котором будем искать значения
    arr := make([]int,100)
    for i := range arr {
        arr[i] = i+1
    }
    //fmt.Println(arr)


    var x int
    fmt.Println("Enter value for search")
    fmt.Scan(&x)

    y := bsearch(arr,x)
    if y != -1 {
        fmt.Printf("%d on position %d\n",x,y)
    } else {
        fmt.Println("no such value")
    }
}

// функция бинарного поиска возвращает -1 если значение не найдено 
// иначе возвращает позицию этого числа в массиве
func bsearch(arr []int, x int) int {
    left, right := 0, len(arr)

    //оба варианта цикла будут работать

    //for left <  right {
    for left != right {
        fmt.Printf("left: %d, right: %d\n",left,right) // чтобы видеть как изменяются left и rigth

        // tmp - это середина текущего промежутка поиска
        tmp := left+(right-left)/2
        if x > arr[tmp] { // если наше число больше числа в серединe массива, 
                          // значит "передвигаем" левую границу на середину поиска и добавляем 1, 
                          // т.к. наше значение точно больше чем значение в середине массива
            left = tmp + 1
            if left == len(arr) { return -1 } // без этого будет выход за пределы массива
        } else {          // если число меньше середины массива, изменяем правую границу
            right = tmp   // не добавляем единицу, т.к. x может равняться arr[right]
        }
        if x == arr[left] {fmt.Printf("FOUND!\n"); return left} // если нашли число, возвращаем его позицию
    }
    // если не нашли значение возвращаем -1
    return -1
}
