package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"кино",
		"кони",
		"Пятак",
		"пятка",
		"слиток",
		"листок",
		"столик",
		"листок",
		"Порт",
		"рог",
		"тяпка"}
	res := getAnagrams(words)
	fmt.Println(res)
}

// Поиск анаграм в слайсе
func getAnagrams(input []string) map[string][]string {
	temp := make(map[string][]string) // Временная мапа с ключом в виде отсортированного по буквам слова
	res := make(map[string][]string)  // Результат
	set := make(map[string]struct{})  // Множество всех слов
	for _, word := range input {
		wordInLower := toLower(word)
		uniqueKey := sortLetters(wordInLower)
		_, isInSet := set[wordInLower] // Проверяем, добавляли ли уже это слово
		if !isInSet {
			temp[uniqueKey] = append(temp[uniqueKey], wordInLower)
			set[wordInLower] = struct{}{}
		}
	}
	for _, v := range temp {
		if len(v) > 1 {
			res[v[0]] = v
			sort.Strings(res[v[0]])
		}
	}
	return res
}

// Приведение к lower case
func toLower(word string) string {
	var res strings.Builder
	for _, r := range word {
		res.WriteRune(unicode.ToLower(r))
	}
	return res.String()
}

// Сортировка букв в строке
func sortLetters(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
