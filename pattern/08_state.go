package pattern

/*
   Паттерн "состояние" позволяет объекту изменять свое поведение в зависимости от его
   внутреннего состояния.

   Плюсы:
   1) Делает код более читаемым и поддерживаемым
   2) Облегчает добавление новых состояний
   3) Позволяет объекту менять свое поведения в рантайме

   Минусы:
   1) Может привести к большому количеству структур, если у объекта много состояний
   2) Может усложнить код

   Примеры использования:
   1) Состояние заказа
   2) Состояние сетевого подключения
*/

type State interface {
	Handle() string
}

type NormalState struct{}

func (s *NormalState) Handle() string {
	return "normal state"
}

type EmergencyState struct{}

func (s *EmergencyState) Handle() string {
	return "emergency state"
}

type Context struct {
	st State
}

func (c *Context) SetState(st State) {
	c.st = st
}

func (c *Context) RequestState() string {
	return c.st.Handle()
}

func main() {
	ctx := &Context{}
	ctx.SetState(&NormalState{})
	ctx.SetState(&EmergencyState{})
}
