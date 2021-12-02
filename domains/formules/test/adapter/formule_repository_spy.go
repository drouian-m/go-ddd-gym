package test_adapter

import (
	"ziggornif.xyz/ddd-gym/domains/formules/model"
	"ziggornif.xyz/ddd-gym/domains/formules/port"
)

type spyRepository struct {
	formules map[string]model.Formule
}

func NewSpyRepository() port.FormuleRepository {
	return &spyRepository{formules: make(map[string]model.Formule)}
}

func (spy *spyRepository) Persist(formule model.Formule) {
	spy.formules[formule.ID] = formule
}

func (spy *spyRepository) GetFormule(ID string) model.Formule {
	return spy.formules[ID]
}
