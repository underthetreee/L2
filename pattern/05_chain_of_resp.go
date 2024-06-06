package pattern

/*
   Паттерн "цепочка вызовов" позволяет передавать запросы последователь через цепочку обработчиков.
   В цепочке каждый обработчик решает может ли он обработать запрос и передает его следующему обработчику
   если не может.

   Плюсы:
   1) Уменьшение зависимости между отправителем запроса и получателем
   2) Упрощение добавление новых обработчиков в цепочку
   3) Позволяет изменять порядок и состав обработчиков

   Минусы:
   1) Возможность сложности отладки какой обработчик обработал запрос
   2) Может привести к созданию длинной цепочки, что усложнит понимание кода

   Примеры использования:
   1) Валидация данных
   2) Логгирование
   3) Обработка HTTP-запросов
*/

type Handler interface {
	Handle(request string) bool
	Next(handler Handler) Handler
}

type FirstHandler struct {
	next Handler
}

func (h *FirstHandler) Next(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *FirstHandler) Handler(request string) bool {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return false
}

type SecondHandler struct {
	next Handler
}

func (h *SecondHandler) Next(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *SecondHandler) Handle(request string) bool {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return false
}
func main() {
	f := &FirstHandler{}
	s := &SecondHandler{}

	f.Next(s)
}
