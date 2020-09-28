package service

import (
	"context"

	"github.com/guspanc/go-crud-pets-api/cmd/api/data"
	"github.com/guspanc/go-crud-pets-api/cmd/api/models"
)

// PetService struct
type PetService struct {
	db data.Datastore
}

// NewPetService returns new pet service
func NewPetService(db data.Datastore) *PetService {
	return &PetService{
		db: db,
	}
}

// GetPets gets all pets
func (ps *PetService) GetPets(ctx context.Context) ([]models.Pet, error) {
	return ps.db.GetPets(ctx)
}

// AddPet adds a pet
func (ps *PetService) AddPet(ctx context.Context, pet models.Pet) (models.Pet, error) {
	if pet.Name == "" {
		return pet, models.NewAPIError("BadRequest", "pet name is required")
	}
	return ps.db.AddPet(ctx, pet)
}
