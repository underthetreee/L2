package pattern

/*
   Паттерн "строитель" используется для создания сложных объектов пошагово.
   Позволяет создавать объекты различной структуры и конфигурации, изолируя
   процесс создания объекта от представления.

   Плюсы:
   1) Позволяет разделить процесс конструирования и представления сложного объекта
   2) Позволяет изменять внутренню структуру создаваемого объекта, не меняя представление
   3) Повышение читаемости кода, путем разделения создания объектов на отдельные шаги

   Минусы:
   1) В случае создания простых объектов может привести к избыточности кода
   2) Переизбыток структур, так как для каждого типа объекта необходим свой строитель

   Примеры использования:
   1) Построение HTTP-запроса
   2) Создание конфигурации для приложения

*/

type Builder interface {
	SetName(name string)
	SetComponentA(name string)
	SetComponentB(name string)
	Build() Product
}

type Product struct {
	name       string
	componentA string
	componentB string
}

type ConcreteBuilder struct {
	p Product
}

func (b *ConcreteBuilder) SetName(name string) {
	b.p.name = name
}

func (b *ConcreteBuilder) SetComponentA(component string) {
	b.p.componentA = component
}

func (b *ConcreteBuilder) SetComponentB(component string) {
	b.p.componentB = component
}

func (b *ConcreteBuilder) Build() Product {
	return b.p
}

type Director struct {
	b Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		b: b,
	}
}

func (d *Director) Construct() Product {
	d.b.SetName("name")
	d.b.SetComponentA("componentA")
	d.b.SetComponentB("componentB")
	return d.b.Build()
}

func main() {
	cb := &ConcreteBuilder{}
	d := NewDirector(cb)
	p := d.Construct()
}
