package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sortFile("input.txt")
}

func sortFile(fileName string) error {
	reverseFlag := flag.Bool("r", false, "sort in reverse order")
	uniqueFlag := flag.Bool("u", false, "delete repeatable strings")
	kFlag := flag.Int("k", 1, "column to order by")
	nFlag := flag.Bool("n", false, "sort by numbers")
	flag.Parse()

	inFile, err := os.Open(fileName)
	var res []string
	if err != nil {
		return err
	}
	defer inFile.Close()

	stringSet := make(map[string]struct{})
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		scannedString := scanner.Text()
		if *uniqueFlag {
			if _, found := stringSet[scannedString]; !found {
				res = append(res, scannedString)
			}
			stringSet[scannedString] = struct{}{}
		} else {
			res = append(res, scannedString)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		firstString := strings.Split(res[i], " ")
		secondString := strings.Split(res[j], " ")
		nColumn := *kFlag - 1
		// Если номер колонки больше числа колонок, то сортируем по первой
		if nColumn >= len(firstString) {
			nColumn = 0
		}
		if *nFlag {
			firstNumber, err := strconv.Atoi(firstString[nColumn])
			if err != nil {
				log.Fatalf("could not convert to int")
			}
			secondNumber, err := strconv.Atoi(secondString[nColumn])
			if err != nil {
				log.Fatalf("could not convert to int")
			}
			return firstNumber < secondNumber
		}
		return firstString[nColumn] < secondString[nColumn]
	})
	outFile, err := os.OpenFile("output.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer outFile.Close()
	if *reverseFlag {
		for i := len(res) - 1; i >= 0; i-- {
			fmt.Fprintf(outFile, "%s\n", res[i])
		}
		return nil
	}
	for _, word := range res {
		fmt.Fprintf(outFile, "%s\n", word)
	}
	return nil
}
