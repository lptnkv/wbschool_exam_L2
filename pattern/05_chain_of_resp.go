package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Поведенческий паттерн уровня объекта
	Есть цепочка получателей, которым передается запрос
	Если запрос не обработан - передается дальше
	Если обработан - может передаваться дальше
	Если не обработан ни одним - теряется
	вместо хранения ссылок на всех кандидатов-получателей запроса, каждый отправитель хранит единственную ссылку на начало цепочки,
	а каждый получатель имеет единственную ссылку на своего преемника - последующий элемент в цепочке.
*/

// Интерфейс обработчика
type Handler interface {
	SendRequest(message int) string
}

// Конкретный обработчик A
type ConcreteHandlerA struct {
	next Handler
}

// Имплементация метода отправки Запроса обработчиком A
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// Конкретный обработчик B
type ConcreteHandlerB struct {
	next Handler
}

// Имплементация метода отправки Запроса обработчиком B
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// Конкретный обработчик C
type ConcreteHandlerC struct {
	next Handler
}

// Имплементация метода отправки Запроса обработчиком C
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
