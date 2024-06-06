package pattern

/*
   Паттерн "посетитель" позволяет добавлять новые операции к объектам без изменения самих объектов.

   Плюсы:
   1) Упрощает добавление новых операций к существующим структурам данных, не изменяя сами структуры
   Минусы:
   2) Внедрение паттерна может сделать код менее понятным и усложнить его структуру.

   Примеры использования:
   1) Анализ графов
   2) Работа с DOM
   3) Обработка AST
*/

type Shape interface {
	Accept()
}

type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type Square struct {
	Width float64
}

func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}

type AreaVisitor struct {
	TotalArea float64
}

func (v *AreaVisitor) VisitCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	v.TotalArea += area
}

func (v *AreaVisitor) VisitSquare(s *Square) {
	area := s.Width + s.Width
	v.TotalArea += area

}

type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
}

func main() {
	c := &Circle{Radius: 3}
	s := &Square{Width: 4}

	av := &AreaVisitor{}

	c.Accept(av)
	s.Accept(av)
}
