package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"strings"
)

// Конструктор структуры Человек
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Структура человек, является фасадом для подсистем Дом, Дерево и Сын
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Метод для возврата списка обязанностей мужчины
func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

// Подсистема Дом
type House struct {
}

// Метод постройки дома
func (h *House) Build() string {
	return "Build house"
}

// Подсистема Дерево
type Tree struct {
}

// Метод для выращивания дерева
func (t *Tree) Grow() string {
	return "Tree grow"
}

// Подсистема Сын
type Child struct {
}

// Метод для рождения сына
func (c *Child) Born() string {
	return "Child born"
}
