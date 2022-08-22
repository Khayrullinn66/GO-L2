package main

import "fmt"

/*
Шаблон фасад (англ. Facade) — структурный шаблон проектирования,
позволяющий скрыть сложность системы путём сведения всех возможных внешних вызовов к одному объекту,
делегирующему их соответствующим объектам системы
Плюсы:
- Изоляция клиентов от поведения сложной системы
Минусы:
- Фасад рискует стать суперклассом, привязанным ко всем классам программы.
Варианты использования:
- Вашему коду приходится работать с большим количеством объектов некой сложной
библиотеки или фреймворка. Вы должны самостоятельно инициализировать эти объекты,
следить за правильным порядком зависимостей и так далее.
В результате бизнес-логика ваших классов тесно переплетается с деталями реализации
сторонних классов. Такой код довольно сложно понимать и поддерживать.
*/

type Cat struct {
}

func (c *Cat) MakeMeow() string {
	return "Meow"
}

type Horse struct {
}

func (h *Horse) MakeNeigh() string {
	return "Neigh"
}

type Dog struct {
}

func (d *Dog) MakeWoof() string {
	return "Woof"
}

type Zoo struct {
	cat   *Cat
	horse *Horse
	dog   *Dog
}

func NewZoo() *Zoo {
	return &Zoo{cat: &Cat{},
		dog:   &Dog{},
		horse: &Horse{}}
}
func main() {
	z := NewZoo()
	fmt.Println(z.cat.MakeMeow())
	fmt.Println(z.dog.MakeWoof())
	fmt.Println(z.horse.MakeNeigh())
}
