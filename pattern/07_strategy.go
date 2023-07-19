package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Стратегия сортировки среза чисел - интерфейс для других сортировок
type StrategySort interface {
	Sort([]int)
}

// Сортировка пузырьком
type BubbleSort struct {
}

// Алгоритм сортировки пузырьком
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// Сортировка вставками
type InsertionSort struct {
}

// Реализация алгоритма сортировки вставками
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// Контекст для выполнения стратегии
type Context struct {
	strategy StrategySort
}

// Выбор стратегии в контексте
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Вызов выбранной стратегии сортировки
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}
