package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	countFlag      bool
	fixedFlag      bool
	ignoreCaseFlag bool
	afterFlag      int
	beforeFlag     int
	contextFlag    int
	invertFlag     bool
	lineNumberFlag bool
)

func init() {
	flag.BoolVar(&countFlag, "c", false, "count strings")
	flag.BoolVar(&fixedFlag, "F", false, "fixed match")
	flag.BoolVar(&ignoreCaseFlag, "i", false, "ignore case")
	flag.IntVar(&afterFlag, "A", 0, "number of strings after match")
	flag.IntVar(&beforeFlag, "B", 0, "number of strings before match")
	flag.IntVar(&contextFlag, "C", 0, "number of lines before and after match")
	flag.BoolVar(&invertFlag, "v", false, "find everything except matches")
	flag.BoolVar(&lineNumberFlag, "n", false, "print string number")
}

func main() {
	flag.Parse()
	grep()
}

func grep() {
	if flag.NArg() < 2 {
		fmt.Println("Not enough arguments")
		return
	}

	// Получаем паттерн поиска
	pattern := flag.Arg(0)

	// Приводим к нижнему регистру, если надо
	if ignoreCaseFlag {
		pattern = strings.ToLower(pattern)
	}

	// Компилируем регулярное выражение
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Could not compile regex: %v\n", err)
	}

	// Открываем заданный файл для чтения
	fileName := flag.Arg(1)
	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Could not open file:  %v", err)
	}
	defer inFile.Close()

	// Счетчик совпадений
	cnt := 0

	// Приводим флаги A и B к единому значению с флагом C
	if contextFlag > beforeFlag {
		beforeFlag = contextFlag
	}
	if contextFlag > afterFlag {
		afterFlag = contextFlag
	}

	// Все считанные строки
	var data []string

	// Сканируем входной файл и заносим в слайс всех строк
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	// Обрабатываем строки слайса
	for lineNumber, line := range data {
		// Если игнорируем регистр, то приводим к нижнему
		if ignoreCaseFlag {
			line = strings.ToLower(line)
		}
		match := false
		// Если нужно точное совпадение со строкой
		if fixedFlag && line == pattern {
			match = true
		} else if r.MatchString(line) {
			match = true
		}

		// Если нужно найти несовпадения
		if invertFlag {
			match = !match
		}

		// Если совпадение найдено, то выводим его при необходимости
		if match {
			cnt++
			// Если не нужно вывести только кол-во совпадений
			if !countFlag {

				// Стартовый индекс в слайсе для вывода строк перед найденной
				start := lineNumber - beforeFlag
				if start < 0 {
					start = 0
				}

				// Выводим строки до найденнрй
				for start < lineNumber {
					fmt.Println(data[start])
					start++
				}

				// Выводим номер строки, если задан флаг
				if lineNumberFlag {
					fmt.Printf("%d:", lineNumber+1)
				}

				// Вывод самой строки
				fmt.Println(line)

				// Индекс последней строки, которую надо вывести после заданной
				end := lineNumber + afterFlag
				if end >= len(data) {
					end = len(data) - 1
				}

				// Вывод строк после найденной
				for i := lineNumber + 1; i < end; i++ {
					fmt.Println(data[i])
				}
			}
		}
	}
	if countFlag {
		fmt.Println(cnt)
	}
}
