package pattern

import "fmt"

/*
   Паттерн "фасад" используется для объединения нескольких интерфейсов подсистемы в один общий,
   тем самым упрощая использование подсистемы, скрывая ее сложность и детали реализации.

   Плюсы:
   1) Сокрытие сложности системы
   2) Упрощение использования
   3) Улучшение читаемости и поддерживаемости
   4) Повышает переиспользуемость кода

   Минусы:
   1) Ограничение гибкости
   2) Возможный излишний уровень абстракции
   3) При увеличении объема фасада, может затруднить поддержку и понимание

   Примеры использования:
   1) API-интерфейсы
   2) Управление транзакциями в БД
   3) Кеширование данных
*/

type ComponentA struct{}

func (c *ComponentA) MethodA() {
	fmt.Println("component A method")
}

type ComponentB struct{}

func (c *ComponentB) MethodB() {
	fmt.Println("component B method")
}

type Facade struct {
	a ComponentA
	b ComponentB
}

func NewFacade() *Facade {
	return &Facade{
		a: ComponentA{},
		b: ComponentB{},
	}
}

func (f *Facade) CallAB() {
	f.a.MethodA()
	f.b.MethodB()
}

func main() {
	f := NewFacade()
	f.CallAB()
}
