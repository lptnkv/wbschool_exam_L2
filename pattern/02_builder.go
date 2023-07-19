package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн Строитель используется, когда нужный продукт сложный и требует нескольких шагов для построения.
В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.
При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного нестабильного
продукта. Паттерн "Строитель" скрывает объект до тех пор, пока он не построен до конца.
*/

// Строитель
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Директор, который управляет строителем
type Director struct {
	builder Builder
}

// Метод постройки, директор говорит строителю, как строить
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// Конкретный строитель
type ConcreteBuilder struct {
	product *Product
}

// Метод конкретного строителя для "постройки" заголовка документа
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

// Метод "постройки" тела документа
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

// Метод "постройки" футера документа
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

// Документ
type Product struct {
	Content string
}

// Возвращает содержимое документа
func (p *Product) Show() string {
	return p.Content
}
