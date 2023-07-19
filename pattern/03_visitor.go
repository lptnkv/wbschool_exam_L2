package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Интерфейс Посетитель
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Интерфейс Место, которое можно посетить
type Place interface {
	Accept(v Visitor) string
}

// Человек, реализует интерфейс Посетитель
type Human struct {
}

// Метод посещения суши-кафе человеком
func (v *Human) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// Суши-кафе реализует интерфейс Место
type SushiBar struct {
}

// Реализация метода Accept интерфейса Place для суши-кафе
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// Внутренний метод структуры Суши-кафе
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Пиццерия реализует интерфейс Место
type Pizzeria struct {
}

// Метод посещения человеком пиццерии
func (v *Human) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// Реализация метода Accept интерфейса Place для пиццерии
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// Внутренний метод структуры Пиццерия, вызывается из метода Accept посетителем
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// Бургерная реализует интерфейс Место
type BurgerBar struct {
}

// Метод посещения человеком бургерной
func (v *Human) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// Реализация метода Accept интерфейса Place для бургерной
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// Внутренний метод структуры Бургерная, вызывается из метода Accept посетителем
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

// Город, имеет коллекцию мест, которые можно посетить
type City struct {
	places []Place
}

// Метод добавления места в коллекцию
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Метод обхода всех Мест в Городе Посетителем
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}
