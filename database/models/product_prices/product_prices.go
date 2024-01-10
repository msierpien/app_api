package model

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type ProductPrices struct {
	ID 					uuid.UUID `json:"id"`
	ProductSuppliersId 	uuid.UUID `json:"product_suppliers_id"`
	Prices 				float64  `json:"prices"`
	EffectiveDate 		time.Time `json:"effective_date"`
	CoedePrices 		float64 `json:"coede_prices"`
	SumPrice 			float64 `json:"sum_price"`
	History 			json.RawMessage `json:"history"` // JSONB
}

func InsertProductPrices(db *sql.DB, pp *ProductPrices) (uuid.UUID, error) {
	query := `INSERT INTO product_prices (product_suppliers_id, prices, effective_date, coede_prices, sum_price, history) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var newID uuid.UUID
	err := db.QueryRow(query, pp.ProductSuppliersId, pp.Prices, pp.EffectiveDate, pp.CoedePrices, pp.SumPrice, pp.History).Scan(&newID)
	if err != nil {
		log.Printf("Błąd podczas wstawiania ProductPrices: %v", err)
		return uuid.Nil, err
	}
	return newID, nil
}

