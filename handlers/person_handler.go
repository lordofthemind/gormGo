package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGo/repositories"
	"github.com/lordofthemind/gormGo/types"
)

type PersonHandler struct {
	Repository repositories.PersonRepository
}

func NewPersonHandler(repo repositories.PersonRepository) *PersonHandler {
	return &PersonHandler{Repository: repo}
}

func (h *PersonHandler) CreatePersonHandler(c *gin.Context) {
	var person types.PersonType
	if err := c.ShouldBindJSON(&person); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, "Failed to bind JSON: "+err.Error())
		return
	}

	createdPerson, err := h.Repository.CreatePerson(&person)
	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, "Failed to create person: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdPerson)
}

func (h *PersonHandler) GetPersonByIdHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handleErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	person, err := h.Repository.GetPersonById(uint(id))
	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, "Failed to get person by ID: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) GetAllPersonsHandler(c *gin.Context) {
	persons, err := h.Repository.GetAllPersons()
	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, "Failed to get all persons: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) UpdatePersonHandler(c *gin.Context) {
	var person types.PersonType
	if err := c.ShouldBindJSON(&person); err != nil {
		handleErrorResponse(c, http.StatusBadRequest, "Failed to bind JSON: "+err.Error())
		return
	}

	updatedPerson, err := h.Repository.UpdatePerson(&person)
	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, "Failed to update person: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedPerson)
}

func (h *PersonHandler) DeletePersonByIdHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handleErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.Repository.DeletePersonById(uint(id))
	if err != nil {
		handleErrorResponse(c, http.StatusInternalServerError, "Failed to delete person by ID: "+err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func handleErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
