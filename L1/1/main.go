package main

import "fmt"

func main() {

	alex := Action{Human: Human{
		Name:  "Alex",
		Age:   25,
		Phone: "+71234567890",
	}}

	alex.SayHi()
	alex.HowOld()
	alex.HaveBirthday()
	alex.HowOld()
	alex.ChangePhone("+70987654321")
	alex.SayHi()

}

// Human
/*
L1.1
Встраивание структур
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human
(аналог наследования). Подсказка: используйте композицию (embedded struct),
чтобы Action имел все методы Human.
*/
type Human struct {
	Name  string
	Age   int
	Phone string
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}
func (h Human) HowOld() {
	fmt.Printf("I`m %d old\n", h.Age)
}
func (h *Human) HaveBirthday() {
	h.Age++
}
func (h *Human) ChangePhone(newPhone string) {
	h.Phone = newPhone
}

type Action struct {
	Human
}
