package port

import "ziggornif.xyz/ddd-gym/domains/formules/model"

type Dispatcher interface {
	Emit(event model.Event)
	//GetEvent(id string) model.FormuleCreatedEvent
}
