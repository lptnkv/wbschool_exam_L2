package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Абстрактное оружие
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
	shoot() string
}

// Конкретное оружие - пушка
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

func (g *Gun) shoot() string {
	return "shot from simple gun"
}

// Конкретное оружие - мушкет
type Musket struct {
	Gun
}

func (musket *Musket) shoot() string {
	return "shot from musket"
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// Конкретное оружие - автомат AK47
type Ak47 struct {
	Gun
}

func (ak *Ak47) shoot() string {
	return "shot from ak47"
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Интерфейс фабрики
type GunFactory struct{}

// Метод создания оружия
func (f *GunFactory) CreateGun(gunType string) IGun {
	switch gunType {
	case "musket":
		return newMusket()
	case "ak47":
		return newAk47()
	default:
		return &Gun{}
	}
}
