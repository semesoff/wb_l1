package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов). 
Реализовать встраивание методов в структуре Action от родительской
 структуры Human (аналог наследования).
*/

// Структура Human
type Human struct {
	Name  string
	Phone string
	Email string
}

// Конструктор структуры Human
func NewHuman(name, phone, email string) *Human {
	return &Human{
		Name:  name,
		Phone: phone,
		Email: email,
	}
}

// Методы структуры Human для изменения имени
func (h *Human) SetName(name string) {
	h.Name = name
}

// Методы структуры Human для изменения телефона
func (h *Human) SetPhone(phone string) {
	h.Phone = phone
}

// Методы структуры Human для изменения почты
func (h *Human) SetEmail(email string) {
	h.Email = email
}

// Методы структуры Human для получения имени
func (h *Human) GetName() string {
	return "Name Human: " + h.Name
}

// Структура Action
type Action struct {
	*Human
}

// Конструктор структуры Action
func NewAction(h *Human) *Action {
	return &Action{
		Human: h,
	}
}

// Методы структуры Action для изменения имени
func (a *Action) PrintName(name string) {
	fmt.Println(a.Human.Name)
}

// Методы структуры Action для изменения телефона
func (a *Action) PrintPhone(phone string) {
	fmt.Println(a.Human.Phone)
}

// Методы структуры Action для изменения почты
func (a *Action) PrintEmail(email string) {
	fmt.Println(a.Human.Email)
}

// Методы структуры Action для получения имени
func (h *Action) GetName() string {
	return "Name Action: " + h.Name
}

func main() {
	human := NewHuman("John", "123456789", "elon@gmail.com")
	action := NewAction(human)

	// Используем методы структуры Human т.к. структура Action включает в себя структуру Human
	action.SetName("Elon")
	action.SetEmail("ggwp@gmail.com")
	action.SetPhone("987654321")

	// Если методы структуры Action имеют одинаковые имена с методами структуры Human,
	// то методы структуры Action будут вызываться т.к. обрабатывается слой вложенности который находится выше
	fmt.Println(action.GetName())
	fmt.Println(action.Human.GetName())
}
