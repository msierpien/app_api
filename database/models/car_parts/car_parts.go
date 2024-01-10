package modelCarParts

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/lib/pq"

	"github.com/google/uuid"
)

// CarPart reprezentuje strukturę rekordu w tabeli car_parts
type CarPart struct {
    IdCarPart    uuid.UUID       `json:"id_car_part"`
    Side         json.RawMessage `json:"side"` // JSONB
    IcIndex      string          `json:"ic_index"`
    TecDocArt    string          `json:"tec_doc_art"`
    TecDocBrand  string          `json:"tec_doc_brand"`
    TowKod       string          `json:"tow_kod"`
    History      json.RawMessage `json:"history"` // JSONB
    CreatedAt    time.Time       `json:"created_at"`
    UpdatedAt    time.Time       `json:"updated_at"`
    // Pola z tabeli products
    ID           uuid.UUID       `json:"id"`
    Name         json.RawMessage `json:"name"` // JSONB
    Slug         json.RawMessage `json:"slug"` // JSONB
    EAN          []string        `json:"ean"`  // text[]
    SKU          string          `json:"sku"`
    Index        string          `json:"index"`
    Status       string          `json:"status"`
    Description  json.RawMessage `json:"description"` // JSONB
    ShortDesc    json.RawMessage `json:"short_description"` // JSONB
    Price        float64         `json:"price"`
    Type         string          `json:"type"` // PRODUCT_TYPE
    Visibility   bool            `json:"visibility"`
    Weight       float64         `json:"weight"`
    Length       float64         `json:"length"`
    Width        float64         `json:"width"`
    Height       float64         `json:"height"`
    CodeCE       int             `json:"code_ce"`
    BrandId      uuid.UUID       `json:"brand_id"`
}


// InsertCarPart dodaje nowy rekord CarPart do bazy danych
func InsertCarPart(db *sql.DB, cp *CarPart) (uuid.UUID, error) {
    // Dodaj RETURNING id do zapytania SQL
    query := `INSERT INTO car_parts (side, ic_index, tec_doc_art, tec_doc_brand, tow_kod, history, created_at, updated_at, name, slug, ean, sku, index, status, description, short_description, price, type, visibility, weight, length, width, height, code_ce, brand_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25) RETURNING id_car_part`

    var newID uuid.UUID
    err := db.QueryRow(query, cp.Side, cp.IcIndex, cp.TecDocArt, cp.TecDocBrand, cp.TowKod, cp.History, cp.CreatedAt, cp.UpdatedAt, cp.Name, cp.Slug, pq.Array(cp.EAN), cp.SKU, cp.Index, cp.Status, cp.Description, cp.ShortDesc, cp.Price, cp.Type, cp.Visibility, cp.Weight, cp.Length, cp.Width, cp.Height, cp.CodeCE, cp.BrandId).Scan(&newID)
    if err != nil {
        log.Printf("Błąd podczas wstawiania CarPart: %v", err)
        return uuid.Nil, err
    }
    // log.Printf("Wstawiono nowy rekord CarPart o ID: %v", newID)
    return newID, nil
}


func ExistsByTowKod(db *sql.DB, towKod string) (bool, error) {
    var exists bool
    query := `SELECT EXISTS(SELECT 1 FROM car_parts WHERE tow_kod = $1)`

    err := db.QueryRow(query, towKod).Scan(&exists)
    if err != nil {
        log.Printf("Błąd podczas sprawdzania istnienia CarPart: %v", err)
        return false, err
    }
    // log.Printf("CarPart o kodzie %s istnieje: %v", towKod, exists)

    return exists, nil
}
