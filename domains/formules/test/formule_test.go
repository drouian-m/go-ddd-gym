package test

import (
	"testing"

	testadapter "ziggornif.xyz/ddd-gym/domains/formules/test/adapter"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"ziggornif.xyz/ddd-gym/domains/formules/model"
	"ziggornif.xyz/ddd-gym/domains/formules/usecase"
)

func TestCreateAFormule(t *testing.T) {
	genID, _ := uuid.NewUUID()
	name := "Formule basique"
	price := 10
	expectedFormule := model.Formule{
		ID:    genID.String(),
		Name:  name,
		Price: price,
	}
	//expectEvent := model.FormuleCreatedEvent{
	//	Event: model.Event{
	//		ID:   genID.String(),
	//		Name: "formule-created",
	//	},
	//}

	spyRepository := testadapter.NewSpyRepository()
	spyDispatcher := testadapter.NewSpyDispatcher()
	createFormuleUseCase := usecase.NewCreateFormuleUseCase(spyRepository, spyDispatcher)
	createFormuleUseCase.Execute(genID.String(), name, price)

	createdFormule := spyRepository.GetFormule(genID.String())
	//emittedEvent := spyDispatcher.GetEvent(genID.String())
	assert.Equal(t, expectedFormule.String(), createdFormule.String())
	//assert.Equal(t, expectEvent.String(), emittedEvent.String())
}
