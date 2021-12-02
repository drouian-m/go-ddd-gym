package model

import (
	"fmt"
)

type Event struct {
	ID   string
	Name string
}

type Formule struct {
	ID            string
	Name          string
	Price         int
	FormuleEvents []FormuleCreatedEvent
}

func NewFormule(id string, name string, price int) Formule {
	createdEvent := NewFormuleCreatedEvent(id)
	formule := Formule{
		ID:    id,
		Name:  name,
		Price: price,
	}

	formule.FormuleEvents = append(formule.FormuleEvents, createdEvent)

	return formule
}

func (f *Formule) String() string {
	return fmt.Sprintf("%v %v %v", f.ID, f.Name, f.Price)
}
