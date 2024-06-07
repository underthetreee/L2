package pattern

import "fmt"

/*
   Паттерн "фабричный метод" определяет интерфейс для создания объекта, но делигирует создание структурам
   реализующий интерфейс.

   Плюсы:
   1) Позволяет изменять типы создаваемых объектов
   2) Помогает разделить ответственность между созданием объектов и их использованием
   3) Расширяемость

   Минусы:
   1) Может привести к усложнению структуры
   2) Избыточность в простых случаях

   Примеры использования:
   1) Web-фреймоворки
   2) Базы данных и ORM
   3) Игровые движки
*/

type Vehicle interface {
	Drive()
}

type Car struct {
}

func (c *Car) Drive() {
	fmt.Println("driving car")
}

type Motocycle struct {
}

func (m *Motocycle) Drive() {
	fmt.Println("driving car")
}

type VehicleFactory interface {
	ProduceVehicle() Vehicle
}

type CarFactory struct {
}

type MotocycleFactory struct {
}

func (f *CarFactory) ProduceVehicle() Vehicle {
	return &Car{}
}

func (f *MotocycleFactory) ProduceVehicle() Vehicle {
	return &Motocycle{}
}

func main() {
	cf := &CarFactory{}

	c := cf.ProduceVehicle()
	c.Drive()

	mf := &MotocycleFactory{}

	m := mf.ProduceVehicle()
	m.Drive()
}
