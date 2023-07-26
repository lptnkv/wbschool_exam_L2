Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	// x = 0, если ничего не передать

	// откладываем вызов функции
	defer func() {
		x++
	}()
	x = 1

	// Тут вызывается отложенная функция, которая изменяет именованное возвращаемое значение
	return
}


func anotherTest() int {
	var x int // x = 0
	defer func() {
		// x = 1
		x++
		// x = 2, но anotherTest уже вернула 1
	}()
	x = 1
	// тут вызывается отложенная функция, но anotherTest уже вернула 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
Deferred functions may read and assign to the returning function’s named return values. (https://go.dev/blog/defer-panic-and-recover)
Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns
Отложенные функции могут изменять именованные параметры
 


```
