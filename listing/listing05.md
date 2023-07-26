Что выведет программа? Объяснить вывод программы.

```go
package main

// Какая-то ошибка со строковым содержимым
type customError struct {
	msg string
}

// Реализация метода Error() интерфейса error
func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	// test возвращает nil и customError имплементирует интерфейс error, так что err.data = nil, err.itab = error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Выведет error, так как интерфейс не равен nil (err.data = nil, но err.itab != nil)

```
