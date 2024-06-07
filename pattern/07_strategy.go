package pattern

import "fmt"

/*
   Паттерн "стратегия" позволяет определить семейство алгоритмов, инкапсулировать их
   и делать их взаимозаменяемыми. Позволяет выбрать алгоритм в зависимости от контекста
   использования.

   Плюсы:
   1) Позволяет добавлять новые алгоритмы без изменения текущего кода
   2) Позволяет избежать дублирования кода
   3) Упрощает тестирование, так как можно протестировать каждый алгоритм отдельно

   Минусы:
   1) Усложнение структуры программы из-за большего количества структур
   2) Повышает сложность отладки

   Примеры использования:
   1) Сортировка
   2) Алгоритмы сжатия
   3) Анализ данных
*/

type SortStrategy interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (s *BubbleSort) Sort(arr []int) []int {
	fmt.Println("bubble sort")
	return arr
}

type InsertionSort struct{}

func (s *InsertionSort) Sort(arr []int) []int {
	fmt.Println("insertion sort")
	return arr
}

type SortContext struct {
	s SortStrategy
}

func NewSortContext(s SortStrategy) *SortContext {
	return &SortContext{
		s: s,
	}
}

func (s *SortContext) SetStrategy(strat SortStrategy) {
	s.s = strat
}

func (s *SortContext) Sort(arr []int) []int {
	return s.Sort(arr)
}

func main() {
	arr := []int{1, 2, 3}
	bs := &BubbleSort{}
	is := &InsertionSort{}

	ctx := NewSortContext(bs)
	arr = ctx.Sort(arr)

	ctx.SetStrategy(is)
	arr = ctx.Sort(arr)
}
