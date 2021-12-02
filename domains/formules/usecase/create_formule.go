package usecase

import (
	"ziggornif.xyz/ddd-gym/domains/formules/model"
	"ziggornif.xyz/ddd-gym/domains/formules/port"
)

type CreateFormuleUseCase interface {
	Execute(ID string, name string, price int)
}

type createFormuleUseCase struct {
	repository port.FormuleRepository
	dispatcher port.Dispatcher
}

func NewCreateFormuleUseCase(repository port.FormuleRepository, dispatcher port.Dispatcher) CreateFormuleUseCase {
	return &createFormuleUseCase{
		repository: repository,
		dispatcher: dispatcher,
	}
}

func (useCase *createFormuleUseCase) Execute(id string, name string, price int) {
	formule := model.NewFormule(id, name, price)
	useCase.repository.Persist(formule)
	for _, event := range formule.FormuleEvents {
		useCase.dispatcher.Emit(event.GetEvent())
	}
}
