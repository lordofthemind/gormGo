// repositories/person_repository_impl.go
package repositories

import (
	"fmt"

	"github.com/lordofthemind/gormGo/initializers"
	"github.com/lordofthemind/gormGo/types"
	"github.com/stretchr/testify/mock"
)

type MainDB struct{}

type MockDB struct {
	mock.Mock
}

// Implement methods of Database interface for MainDB.
func (db *MainDB) Create(value interface{}) error {
	return initializers.DB.Create(value).Error
}

func (db *MainDB) First(out interface{}, where ...interface{}) error {
	return initializers.DB.First(out, where...).Error
}

func (db *MainDB) Find(out interface{}, where ...interface{}) error {
	return initializers.DB.Find(out, where...).Error
}

func (db *MainDB) Save(value interface{}) error {
	return initializers.DB.Save(value).Error
}

func (db *MainDB) Delete(value interface{}, where ...interface{}) error {
	return initializers.DB.Delete(value, where...).Error
}

// PersonRepositoryImpl represents the implementation of PersonRepository.
type PersonRepositoryImpl struct {
	DB Database
}

func NewPersonRepository() PersonRepository {
	return &PersonRepositoryImpl{
		DB: &MainDB{},
	}
}

func (r *PersonRepositoryImpl) CreatePerson(person *types.PersonType) (*types.PersonType, error) {
	if err := initializers.DB.Create(person).Error; err != nil {
		return nil, fmt.Errorf("failed to create person: %w", err)
	}
	return person, nil
}

func (r *PersonRepositoryImpl) GetPersonById(id uint) (*types.PersonType, error) {
	person := &types.PersonType{}
	if err := initializers.DB.First(person, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get person by ID %d: %w", id, err)
	}
	return person, nil
}

func (r *PersonRepositoryImpl) GetAllPersons() ([]types.PersonType, error) {
	var persons []types.PersonType
	if err := initializers.DB.Find(&persons).Error; err != nil {
		return nil, fmt.Errorf("failed to get all persons: %w", err)
	}
	return persons, nil
}

func (r *PersonRepositoryImpl) UpdatePerson(person *types.PersonType) (*types.PersonType, error) {
	if err := initializers.DB.Save(person).Error; err != nil {
		return nil, fmt.Errorf("failed to update person with ID %d: %w", person.ID, err)
	}
	return person, nil
}

func (r *PersonRepositoryImpl) DeletePersonById(id uint) error {
	if err := initializers.DB.Delete(&types.PersonType{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete person with ID %d: %w", id, err)
	}
	return nil
}
