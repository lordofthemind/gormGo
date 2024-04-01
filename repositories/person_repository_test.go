// person_repositories_test.go

package repositories_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/lordofthemind/gormGo/helpers"
	"github.com/lordofthemind/gormGo/repositories"
	"github.com/lordofthemind/gormGo/tests"
	"github.com/lordofthemind/gormGo/types"
	"github.com/stretchr/testify/require"
)

func TestCreatePerson(t *testing.T) {
	// Initialize the database connection for testing
	if err := tests.InitializeForTesting(); err != nil {
		log.Fatalf("Failed to initialize test environment: %v", err)
	}
	fmt.Println("Database connection initialized for testing")

	// Ensure that MOCK_DB is properly initialized
	require.NotNil(t, tests.MOCK_DB, "MOCK_DB is nil")
	// Initialize repository
	repo := repositories.NewPersonRepository()

	// Create a random generator
	rg := helpers.NewRandomGenerator()

	person := &types.PersonType{
		// Generate random values for each field
		Username: rg.RandomValue("Username").(string),
		Email:    rg.RandomValue("Email").(string),
		Phone:    rg.RandomValue("Phone").(string),
		Name:     rg.RandomValue("Name").(string),
		Address:  rg.RandomValue("Address").(string),
		Age:      rg.RandomValue("Age").(uint),
		Gender:   rg.RandomValue("Gender").(string),
	}

	// Call repository function
	createdPerson, err := repo.CreatePerson(person)

	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, person.Username, createdPerson.Username)
	require.Equal(t, person.Email, createdPerson.Email)
	require.Equal(t, person.Phone, createdPerson.Phone)
	require.Equal(t, person.Name, createdPerson.Name)
	require.Equal(t, person.Address, createdPerson.Address)
	require.Equal(t, person.Age, createdPerson.Age)
	require.Equal(t, person.Gender, createdPerson.Gender)
	require.NotZero(t, createdPerson.ID)
	require.NotZero(t, createdPerson.CreatedAt)
}

// Add similar tests for other repository functions (GetPersonById, GetAllPersons, UpdatePerson, DeletePersonById)
