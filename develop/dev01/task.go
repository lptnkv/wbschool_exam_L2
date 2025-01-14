package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	currTime, err := getTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Current time: ", currTime)
}

func getTime() (string, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return "", err
	}
	return time.String(), nil
}
