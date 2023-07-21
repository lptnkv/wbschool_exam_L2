package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	res, err := rleDecode("abcd")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(res)
	}

	res, err = rleDecode("")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(res)
	}

	res, err = rleDecode("45")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(res)
	}

	res, err = rleDecode("a5")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(res)
	}

	res, err = rleDecode("a5b3ch1")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(res)
	}

	res, err = rleDecode("g20")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(res)
	}
}

func rleDecode(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	runes := []rune(input)
	if unicode.IsDigit(runes[0]) {
		return "", fmt.Errorf("incorrect string")
	}
	var result strings.Builder
	length := len(runes)
	i := 0
	// Идем по символам строки
	for i < length {
		// Встретили букву
		if unicode.IsLetter(runes[i]) {
			// Если следующая цифра и не вышли за границу строки
			if i+1 < length && unicode.IsDigit(runes[i+1]) {
				cnt := int(runes[i+1] - '0')
				symbolToWrite := runes[i]
				i++
				for i+1 < length && unicode.IsDigit(runes[i+1]) {
					i++
					cnt = cnt*10 + int(runes[i]-'0')
				}
				for j := 0; j < cnt; j++ {
					result.WriteRune(symbolToWrite)
				}
			} else {
				result.WriteRune(runes[i])
			}
		}
		i++
	}
	return result.String(), nil
}
