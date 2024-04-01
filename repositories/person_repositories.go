// repositories/person_repository.go
package repositories

import "github.com/lordofthemind/gormGo/types"

type PersonRepository interface {
	CreatePerson(person *types.PersonType) (*types.PersonType, error)
	GetPersonById(id uint) (*types.PersonType, error)
	GetAllPersons() ([]types.PersonType, error)
	UpdatePerson(person *types.PersonType) (*types.PersonType, error)
	DeletePersonById(id uint) error
}
