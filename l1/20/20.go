package main

import (
	"fmt"
)

func main() {
    str := "    sun   dog   snow      "
    fmt.Printf("string: %s\n", str)
    wordsSlice := words(str)
    end := len(wordsSlice)-1
    for i := 0; i <= end/2; i++ {
        wordsSlice[i],wordsSlice[end-i] = wordsSlice[end-i],wordsSlice[i]
    }
    fmt.Printf("reversed string: %s\n", unwords(wordsSlice))
}


// функция words "разбивает" строку на отдельные слова. Возвращает слайс этих слов.
// Слова в строке могут быть разделены произвольным количеством пробелов
func words(s string) []string {
    var res []string
    
    //удаляем пробелы в конце строки если они есть
    for s[len(s)-1] == ' ' {
        s = s[:len(s)-1]
    }

    //разбиваем строку на слова
    var start int = 0     // первый байт слова
    var b bool            // для определения конца слова
    for i := range s {
        if s[i] != ' ' {  // если нашли НЕ пробел
            if !b {       // и у нас не было найдено символов
                b = true  // нашли символ
                start = i // отсюда начинается слово
            }
        } else if b {     // если символ это пробел и у нас был найден символ
                          // значит сейчас слово закончилось
            res = append(res,s[start:i]) // добавляем слово
            start = i + 1 // cледующее слово начинается как минимум со следующего символа
            b = false     // все найденные символы записаны в слово,
        }
    }
    // добавляем последнее слово
    res = append(res,s[start:])
    return res
}

// объединяет слайс строк в одну строку
func unwords(s []string) string {

    var size int = 0 // длина итоговой строки в байтах
    for i := range s {
        size += len(s[i])
    }
    size += len(s)  // не забываем про пробелы

    res := make([]byte, size)

    var pos int = 0 // позиция для записи очередного символа
    for i := range s {
        for j := 0; j < len(s[i]); j++ {
            res[pos] = s[i][j]
            pos++
        }
        res[pos] = ' '
        pos++
    }

    return string(res[:len(res)-1]) //пропускаем пробел в конце строки

}
















