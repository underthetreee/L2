Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
- В функции test() создаётся новый объект типа customError с сообщением "something went wrong", и его указатель возвращается как тип *customError.
- В функции main() переменная err объявлена как error, и в неё присваивается результат вызова test(), который является типом *customError.
- Далее проверяется условие if err != nil. Поскольку err является указателем на объект типа customError, который был возвращён из test(), он не равен nil. Таким образом, условие if err != nil выполняется.
