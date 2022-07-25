package main

import (
    "fmt"
    "math/big"
    "os"
)

type BigInt []byte

func main () {
    var a,b,res big.Int
    //var as,bs string
    fmt.Println("enter a")
    fmt.Fscan(os.Stdin,&a)
    fmt.Println("enter b")
    fmt.Fscan(os.Stdin,&b)

    fmt.Printf("a + b: %v\n",res.Add(&a,&b))
    fmt.Printf("a - b: %v\n",res.Sub(&a,&b))
    fmt.Printf("a/b: %v\n",res.Div(&a,&b))
    fmt.Printf("a*b: %v\n",res.Mul(&a,&b))

    /*fmt.Println("my work:")
    var x,y BigInt
    fmt.Println("enter x")
    fmt.Fscan(os.Stdin,&x)
    fmt.Println("enter y")
    fmt.Fscan(os.Stdin,&y)
    fmt.Printf("x + y: %s\n",BigSum(x,y))
    fmt.Printf("x - y: %s\n",BigSub(x,y))*/
}

// собственная реализация математики больших чисел
// реализованы только сумма и разность
func BigSum(x,y BigInt) BigInt {

    var res []byte
    // предпологаем что x длиннее y 
    bigLen   := len(x)
    smallLen := len(y)

    // чтобы избежать многочисленных if'ов
    bigVal   := x

    // проверяем предположение
    if len(y) > len(x) { // если предположение не верно меняем местами значения
        bigLen, smallLen = len(y), len(x)
        bigVal = y
    }

    // проверка на возможность переполнения
    if len(x) == len(y) { // если длины чисел равны
        if x[0]-'0' + y[0]-'0' > 8 { // переполнение возможно только если сумма
                                     // первых элементов равна 9
            // увеличиваем длинну итогового слайса на 1 (для переполнения)
            res = make([]byte, bigLen + 1) 
        } else {
            res = make([]byte, bigLen)
        }
    } else { // длины не равны
        // переполнение возможно только если первый элемент длинного слайса '9'
        if bigVal[0] > 8 + '0' {
            res = make([]byte, bigLen + 1) 
        } else { // переполнения быть не может
            res = make([]byte, bigLen) 
        }
    }

    // если будет переполнение то '0' заменится на 1,
    res[0] = '0'
    over := 0
    i := 0

    for i < smallLen {

        lastx := len(x) -1 -i
        lasty := len(y) -1 -i

        tmp := int(x[lastx]-'0') + int(y[lasty]-'0') + over
        res[len(res) -i -1] = byte(tmp%10 + '0')

        over = int(tmp/10)
        i++
    }

    for i < bigLen {

        lastx := bigLen -1 -i

        tmp := int(x[lastx]-'0') + over
        res[len(x) -i -1] = byte(tmp%10 + '0')

        over = int(tmp/10)
        i++
    }

    // если есть переполнение 
    if over != 0 { res[0] = byte(1+'0') }

    // если резервировали ячейку под переполнение, но его не произошло
    if res[0] == '0' { 
        // "сдвигаем" все числа на 1 влево
        for i := 0; i < len(res)-1; i++ {
            res[i] = res[i+1]
        }
        // после сдвига игнорируем последнее значение
        return res[:len(res)-1]
    }

    return res
}

// вычитает y из x
func BigSub(x,y BigInt) BigInt {
    var res []byte
    // если вычитаем отрицательное из отрицательного, например -2 (-) (-1)
    if x[0] == '-' && y[0] == '-' {
        tmp_res,neg := trueSub(x[1:],y[1:])
        if neg { // результат вычитания меньше нуля, но у нас уже есть -,
                 // значит итоговое значение будет положительное
            return tmp_res
        } else { // результат положительный, итоговое значение отрицательное
            res := make([]byte, len(tmp_res) + 1)
            res[0] = '-'
            for i := 0; i < len(res)-1; i++ {
                res[len(res)-1-i] = tmp_res[len(tmp_res)-1-i]
            }
            return res
        }
    }

    // если вычитаем положительное из отрицательного, например -5 (-) 2
    if x[0] == '-' && y[0] != '-' {
        tmp_res := BigSum(x[1:],y)
        res := make([]byte, len(tmp_res) + 1)
        res[0] = '-'
        for i := 0; i < len(tmp_res)-1; i++ {
            res[len(res)-1-i] = tmp_res[len(tmp_res)-1-i]
        }
        return res
    }

    // если вычитаем отрицательное из положительного, например 4 (-) -2
    if x[0] != '-' && y[0] == '-' { 
        return BigSum(x,y[1:])
    }

    // если вычитаем положительное из положительного, например 6 (-) 3
    if x[0] != '-' && y[0] != '-' {
        tmp_res,neg := trueSub(x,y)
        if neg {
            res := make([]byte, len(tmp_res) + 1)
            res[0] = '-'

            for i := 0; i < len(tmp_res); i++ {
                res[i+1] = tmp_res[i]
            }
        return res
        } else {
            return tmp_res
        }
    }

    fmt.Println("YOU WILL NEVER SEE THIS!!!!")
    return res
}

// вспомогательная функция которая реализует логику вычитания больших чисел
// оба числа положительные 
func trueSub(x,y BigInt) (BigInt, bool) {
    var res []byte
    var neg bool = false

    // для избежания изменения входных данных явно скопируем оба числа
    // предпологаем что x > y

    big := make([]byte,len(x))
    for i := range x {
        big[i] = x[i]
    }
    small := make([]byte,len(y))
    for i := range y {
        small[i] = y[i]
    }
    // если оба числа равны по длинне
    if len(y) == len(x) {
        i := 0
        // найдем первое отличающееся число
        for i < len(x) && x[i] == y[i] {
            if x[i] != y[i] {
                break
            } else {
                i++
            }
            if i == len(x) {
                i--; break
            }
        }

        if x[i] < y[i] { // y > x
            big,small = small, big
            neg   = true
        } else if x[i] == y[i] && i == len(x)-1 {// x = y, при вычитании получим 0
            return []byte{'0'}, false
        } 
        // x > y, ничего не изменяем
    }

    // если вычитаемое длиннее вычитаемого, оно автоматически больше
    if len(y) > len(x) { // y > x
        big,small = small,big
        neg   = true
    }

    bigLen   := len(big)
    smallLen := len(small)

    res = make([]byte, bigLen)
    for i := range res {
        res[i] = '0'
    }

    i := 0
    for i < smallLen {

        bigEnd  := bigLen-1-i
        smallEnd := smallLen-1-i

        if big[bigEnd] == '0' && small[smallEnd] != '0' {
            tmp := getOne(big[:bigEnd]) + int(big[bigEnd])-'0'
            tmp -= int(small[smallEnd]) - '0'
            res[bigEnd] = byte(tmp) + '0'
        } else if big[bigEnd] < small[smallEnd] {
            tmp := getOne(big[:bigEnd]) + int(big[bigEnd])-'0'
            tmp -= int(small[smallEnd]) - '0'
            res[bigEnd] = byte(tmp) + '0'
        } else {
            tmp := big[bigEnd] - small[smallEnd]
            res[bigEnd] = byte(tmp) + '0'
        }

        i++
    }

    for i < bigLen {
        end := bigLen-1-i
        res[end] = big[end]
        i++
    }

    zero := 0
    for res[zero] == '0' {
        zero++
    }
    return res[zero:], neg
}

// вспомогательная функция для вспомогательной функции
// которая "забирает" единичку из старшего разряда.
// если значение этого разряда == 0, она обращается к следующему старшему,
// и так до тех пор пока не найдет нужный разряд,
// после нахождения нужного разряда, все промежуточные разряды(нули)
// заполняет '9'

func getOne(x BigInt) int {

    end := len(x)-1
    start := end
    b := false

    if x[len(x)-1] == '0' {
        end = len(x)-1
        start--
        b = true
        for start >= 0 {
            if x[start] == '0' {
                start--
            } else {
                x[start] -= 1
                start++
                break;
            }
        }
    } else {
        x[len(x)-1] -= 1
        return 10
    }

    if b {
        for start <= end {
            x[start] = '9'
            start++
        }
    }
    return 10
}
