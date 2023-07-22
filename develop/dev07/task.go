package main

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	<-makeDone(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Hour),
		sig(6*time.Minute),
		sig(10*time.Second),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}

// Функция создания done-канала и его закрытия при закрытии одного из переданных done-каналов
func makeDone(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если передано 0 каналов, то возвращаем nil
		return nil
	case 1:
		// Если передан один канал, то возвращаем его
		return channels[0]
	}

	// Итоговый канал
	orDone := make(chan interface{})

	// Запускаем горутину, которая получает сигнал от любого из каналов и затем закрывает итоговый канал
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			// Если передано два канала, то дожидаемся done-сигнала от любого из них
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			// Если передано 3 и более каналов, то рекурсивно объединяем их и ждем сигнала от любого
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-makeDone(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

// Функция для создания канала, который закроется спустя заданное время
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
