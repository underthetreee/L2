package pattern

import "fmt"

/*
   Паттерт "комманда" позволяет инкапсулировать операции в отдельные объекты, позволяя параметризовать клиентов с конкретными запросами,
   откладывать выполнение операций, записывать историю и поддерживать отмену операций

   Плюсы:
   1) Гибкость и расширяемость
   2) Поддержка отмены и отложенного выполнения операций
   3) Отделение ответственности

   Минусы:
   1) Увеличения числа структур
   2) Усложнение отладки из-за дополнительной абстракции

   Примеры использования:
   1) Управление очередью задач
   2) Обработка запросов HTTP-сервером
*/

type Command interface {
	Execute()
}

type FirstCommand struct {
}

func (c *FirstCommand) Execute() {
	fmt.Println("execute first command")
}

type SecondCommand struct {
}

func (c *SecondCommand) Execute() {
	fmt.Println("execute second command")
}

func RunCommand(cmd Command) {
	cmd.Execute()
}

func main() {
	fc := &FirstCommand{}
	sc := &SecondCommand{}

	RunCommand(fc)
	RunCommand(sc)
}
