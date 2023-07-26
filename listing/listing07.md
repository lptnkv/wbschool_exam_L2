Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция принимает переменное кол-во интов, кидает их в канал, закрывает его и возвращает канал
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

// Функция слияния двух каналов
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			// Нужна проверка на закрытие канала: v, ok := <- a; 
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
		// Тут нужно закрыть канал
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа выведет числа от 1 до 8 в случайном порядке (зависит от того, какой канал выбрал select), затем будет выводить нули, потому что в функции merge нет проверки на закрытие каждого канала (v, ok := <- c), а читать из закрытого канала можно (в отличие от nil), но будет значение по умолчанию

```
