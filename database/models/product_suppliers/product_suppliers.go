package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)


type ProductSuppliers struct {
    ID          uuid.UUID       `json:"id"`
    ProductId   uuid.UUID       `json:"product_id"`
	CarPartId   uuid.UUID       `json:"car_part_id"`
    SupplierId  uuid.UUID       `json:"supplier_id"`
    IndexSuppliers string       `json:"index_suppliers"`
    CreatedAt   time.Time       `json:"created_at"`
    UpdatedAt   time.Time       `json:"updated_at"`

}

func InsertCarPartSuppliers(db *sql.DB, ps *ProductSuppliers) error {
	query := `INSERT INTO product_suppliers ( car_part_id, supplier_id, index_suppliers, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(query, ps.CarPartId, ps.SupplierId, ps.IndexSuppliers, ps.CreatedAt, ps.UpdatedAt)
	if err != nil {
		log.Printf("Błąd podczas wstawiania ProductSuppliers: %v", err)
		return err
	}

	return nil 
}