package main

import "fmt"

type Human struct {
	Name  string
	Phone string
	Email string
}

func NewHuman(name, phone, email string) *Human {
	return &Human{
		Name:  name,
		Phone: phone,
		Email: email,
	}
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetPhone(phone string) {
	h.Phone = phone
}

func (h *Human) SetEmail(email string) {
	h.Email = email
}

func (h *Human) GetName() string {
	return "Name Human: " + h.Name
}

type Action struct {
	*Human
}

func NewAction(h *Human) *Action {
	return &Action{
		Human: h,
	}
}

func (a *Action) PrintName(name string) {
	fmt.Println(a.Human.Name)
}

func (a *Action) PrintPhone(phone string) {
	fmt.Println(a.Human.Phone)
}

func (a *Action) PrintEmail(email string) {
	fmt.Println(a.Human.Email)
}

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
