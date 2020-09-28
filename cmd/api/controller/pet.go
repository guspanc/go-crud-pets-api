package controller

import (
	"net/http"

	"github.com/guspanc/go-crud-pets-api/cmd/api/models"
	"github.com/guspanc/go-crud-pets-api/cmd/api/service"
)

// PetController struc
type PetController struct {
	svc *service.PetService
}

// NewPetController returns new pet controller
func NewPetController(svc *service.PetService) *PetController {
	return &PetController{svc}
}

// HandleGetPets request
func (pc *PetController) HandleGetPets(w http.ResponseWriter, r *http.Request) {
	pets, err := pc.svc.GetPets(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	encodeResponse(w, pets)
}

// HandleAddPet request
func (pc *PetController) HandleAddPet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	err := decodeRequest(r, &pet)
	if err != nil {
		handleError(w, err)
		return
	}
	pet, err = pc.svc.AddPet(r.Context(), pet)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	encodeResponse(w, pet)
}
