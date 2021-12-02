package port

import "ziggornif.xyz/ddd-gym/domains/formules/model"

type FormuleRepository interface {
	Persist(formule model.Formule)
	GetFormule(ID string) model.Formule
}
