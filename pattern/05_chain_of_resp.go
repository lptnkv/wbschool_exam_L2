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

// Помещение больницы
type Department interface {
	execute(*Patient) string
	setNext(Department)
}

// Пациент
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Конкретный обработчик - Приемное отделение
type Reception struct {
	next Department
}

// Реализация метода обработки Пациента
func (r *Reception) execute(p *Patient) string {
	res := ""
	if p.registrationDone {
		res += "Patient registration already done\n"
		r.next.execute(p)
		return res
	}
	res += "Reception registering patient\n"
	p.registrationDone = true
	res += r.next.execute(p)
	return res
}

// Метод добавления следующего обработчика
func (r *Reception) setNext(next Department) {
	r.next = next
}

// Доктор
type Doctor struct {
	next Department
}

// Обработчик кабинета Доктора
func (d *Doctor) execute(p *Patient) string {
	res := ""
	if p.doctorCheckUpDone {
		res += "Doctor checkup already done\n"
		d.next.execute(p)
		return res
	}
	res += "Doctor checking patient\n"
	p.doctorCheckUpDone = true
	res += d.next.execute(p)
	return res
}

// Метод добавления следующего обработчика
func (d *Doctor) setNext(next Department) {
	d.next = next
}

// Аптека
type Medical struct {
	next Department
}

// Обработчик Аптеки
func (m *Medical) execute(p *Patient) string {
	res := ""
	if p.medicineDone {
		res += "Medicine already given to patient\n"
		m.next.execute(p)
		return res
	}
	res += "Medical giving medicine to patient\n"
	p.medicineDone = true
	res += m.next.execute(p)
	return res
}

// Метод добавления следующего обработчика после Аптеки
func (m *Medical) setNext(next Department) {
	m.next = next
}

// Касса
type Cashier struct {
	next Department
}

// Обработчик Кассы
func (c *Cashier) execute(p *Patient) string {
	res := ""
	if p.paymentDone {
		res += "Payment Done"
		return res
	}
	res += "Cashier getting money from patient"
	return res
}

// Метод добавления следующего обработчика после Кассы
func (c *Cashier) setNext(next Department) {
	c.next = next
}
