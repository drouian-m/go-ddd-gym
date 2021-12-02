package test_adapter

import (
	"ziggornif.xyz/ddd-gym/domains/formules/model"
	"ziggornif.xyz/ddd-gym/domains/formules/port"
)

type spyDispatcher struct {
	events map[string]model.Event
}

func NewSpyDispatcher() port.Dispatcher {
	return &spyDispatcher{events: make(map[string]model.Event)}
}

func (spy *spyDispatcher) Emit(event model.Event) {
	spy.events[event.ID] = event
}

func (spy *spyDispatcher) GetEvent(id string) model.Event {
	return spy.events[id]
}
