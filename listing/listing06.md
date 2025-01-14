Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"} // len = 3, cap = 3
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3" // изменяет массив, на который ссылается слайс
	i = append(i, "4") // Добавляет новый элемент, len>cap, поэтому выделяется память для нового массива и i ссылается на него
	i[1] = "5" // Меняется элемент в новом массиве
	i = append(i, "6") // Добавляется элемент в новом массиве
	// i = [3, 5, 3, 4, 6]
}
```

Ответ:
```
Вывод: [3 2 3]
В функции modifySlice изменится 0-й элемент массива под слайсом s, но затем при добавлении новых элементов длина выйдет за capacity и выделится новая память, затем элемент с индексом 1 поменяется в новом массиве и элемент "6" так же добавится в новый массив
```

```go
type slice struct {
	array unsafe.Pointer // Указатель на данные
	len   int // Длина - текущее кол-во элементов
	cap   int // Вместимость - максимальное кол-во элементов

}
```

```
при передаче в функцию передается структура SliceHeader
при переполнении слайса (len > cap) выделяется новая память

Рост capacity слайса
starting cap    growth factor
256             2.0
512             1.63
1024            1.44
2048            1.35
4096            1.30

Или по следующей формуле: newcap += (newcap + 3*threshold) / 4, где threshold = 256
Встроенная функция append возвращает новую структуру SliceHeader

```
