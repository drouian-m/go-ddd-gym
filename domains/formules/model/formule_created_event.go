package model

import "fmt"

type FormuleCreatedEvent struct {
	Event
}

func NewFormuleCreatedEvent(id string) FormuleCreatedEvent {
	return FormuleCreatedEvent{
		Event: Event{
			ID:   id,
			Name: "formule-created",
		},
	}
}

func (f *FormuleCreatedEvent) String() string {
	return fmt.Sprintf("%v %v", f.ID, f.Name)
}

func (f *FormuleCreatedEvent) GetEvent() Event {
	return f.Event
}
