package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fieldsFlag    string
	delimiterFlag string
	separatedFlag bool
)

func init() {
	flag.StringVar(&fieldsFlag, "f", "", "choose fields to cut")
	flag.StringVar(&delimiterFlag, "d", "\t", "choose delimiter")
	flag.BoolVar(&separatedFlag, "s", false, "only separated strings")

}

func main() {
	cut()
}

func cut() {
	flag.Parse()

	// Итоговый вывод
	var res []string

	// Парсим флаг -f
	var fields []int
	if fieldsFlag != "" {
		// Если записан диапазон колонок через дефис
		if strings.Contains(fieldsFlag, "-") {
			start, err := strconv.Atoi(strings.Split(fieldsFlag, "-")[0])
			if err != nil {
				fmt.Printf("Can't convert field flag to integer: %v\n", err.Error())
				os.Exit(1)
			}
			end, err := strconv.Atoi(strings.Split(fieldsFlag, "-")[1])
			if err != nil {
				fmt.Printf("Can't convert field flag to integer: %v\n", err.Error())
				os.Exit(1)
			}
			for start <= end {
				fields = append(fields, start-1)
				start++
			}
		} else if strings.Contains(fieldsFlag, ",") {
			// Если номера колонок записаны через запятую, то парсим и
			for _, fieldNumber := range strings.Split(fieldsFlag, ",") {
				convertedFieldNumber, err := strconv.Atoi(fieldNumber)
				if err != nil {
					fmt.Printf("Can't convert field flag to integer: %v\n", err.Error())
					os.Exit(1)
				}
				fields = append(fields, convertedFieldNumber-1)
			}
		} else {
			// Если записано одно число
			fieldNumber, err := strconv.Atoi(fieldsFlag)
			if err != nil {
				fmt.Printf("Can't convert field flag to integer: %v\n", err.Error())
				os.Exit(1)
			}
			fields = append(fields, fieldNumber-1)
		}
	}

	// Сканируем стандартный ввод построчно, в конце следут ctrl-z
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// Если установлен флаг separated и разделителя нет, то пропускаем строку
		if separatedFlag && !strings.Contains(line, delimiterFlag) {
			continue
		}

		// Разделяем строку по разделителю на слова
		words := strings.Split(line, delimiterFlag)
		var resultLine string

		// Если нужно вывести конкретные колонки
		if fieldsFlag != "" {
			// Слайс для хранения необходимых колонок
			var matchedWords []string

			// Соединяем в одну строку и добавляем к результату
			for _, fieldNumber := range fields {
				// Если номер колонки больше кол-ва слов в строке, то пропускаем
				if fieldNumber >= len(words) {
					continue
				}
				matchedWords = append(matchedWords, words[fieldNumber])
			}
			resultLine = strings.Join(matchedWords, delimiterFlag)
		} else {
			resultLine = line
		}
		res = append(res, resultLine)
	}
	for _, line := range res {
		fmt.Println(line)
	}
}
