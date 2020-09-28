package data

import (
	"context"
	"database/sql"

	"github.com/guspanc/go-crud-pets-api/cmd/api/models"

	"github.com/google/uuid"
)

const (
	getPetsQuery = `
		SELECT BIN_TO_UUID(p.id) AS pet_id, p.name, p.breed, BIN_TO_UUID(o.id) AS owner_id, o.fullName
		FROM pets AS p
		LEFT JOIN owners AS o ON p.ownerId = o.id
	`
	addPetQuery = `
		INSERT INTO pets (id, name, breed, ownerId)	
		VALUES (UUID_TO_BIN(?), ?, ?, UUID_TO_BIN(?))
	`
	addOwnerQuery = `
		INSERT INTO owners (id, fullName)
		VALUES (UUID_TO_BIN(?), ?)
	`
)

// GetPets returns all pets
func (db *DB) GetPets(ctx context.Context) ([]models.Pet, error) {
	rows, err := db.QueryContext(ctx, getPetsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	pets := []models.Pet{}
	for rows.Next() {
		pet := models.Pet{}
		var ownerID uuid.UUID
		var ownerFullName sql.NullString
		err = rows.Scan(&pet.ID, &pet.Name, &pet.Breed, &ownerID, &ownerFullName)
		if err != nil {
			return nil, err
		}
		if ownerID != uuid.Nil || ownerFullName.String != "" {
			pet.Owner = &models.Owner{
				ID:       ownerID,
				FullName: ownerFullName.String,
			}
		}
		pets = append(pets, pet)
	}
	return pets, nil
}

// AddPet adds a pet
func (db *DB) AddPet(ctx context.Context, pet models.Pet) (models.Pet, error) {
	// begin tx
	tx, err := db.Begin()
	if err != nil {
		return pet, err
	}
	var ownerID *string
	if pet.Owner != nil && (pet.Owner.ID != uuid.Nil || pet.Owner.FullName != "") {
		if pet.Owner.ID == uuid.Nil {
			// add owner to table
			pet.Owner.ID = uuid.New()
			_, err := tx.ExecContext(ctx, addOwnerQuery, pet.Owner.ID.String(), pet.Owner.FullName)
			if err != nil {
				tx.Rollback()
				return pet, err
			}
		}
		ownerIDStr := pet.Owner.ID.String()
		ownerID = &ownerIDStr
	}
	pet.ID = uuid.New()
	_, err = tx.ExecContext(ctx, addPetQuery, pet.ID.String(), pet.Name, pet.Breed, ownerID)
	if err != nil {
		tx.Rollback()
		return pet, err
	}
	err = tx.Commit()
	return pet, err
}
