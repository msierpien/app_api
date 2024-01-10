package modelBrands

import (
    "database/sql"
    "time"

    "github.com/google/uuid"
)

type Brand struct {
	ID           uuid.UUID       `json:"id"`
	Name 	   	 string          `json:"name"` // JSONB
	CreatedAt    time.Time       `json:"created_at"`
    UpdatedAt    time.Time       `json:"updated_at"`
}

func InsertBrand(db *sql.DB, b *Brand) (uuid.UUID, error) {
    // Dodaj klauzulę RETURNING do zapytania, aby zwrócić ID nowo wstawionego wiersza
    query := `INSERT INTO brand (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

    var newID uuid.UUID
    err := db.QueryRow(query, b.ID, b.Name, b.CreatedAt, b.UpdatedAt).Scan(&newID)
    if err != nil {
        return uuid.Nil, err
    }

    return newID, nil
}

func SelectBrand(db *sql.DB, id string) (*Brand, error) {
	query := `SELECT id, name, created_at, updated_at FROM brand WHERE id = $1`
	row := db.QueryRow(query, id)

	var b Brand
	err := row.Scan(&b.ID, &b.Name, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func GetBrandIDByName(db *sql.DB, name string) (uuid.UUID, error) {
    var id uuid.UUID
    query := `SELECT id FROM brand WHERE name = $1`

    err := db.QueryRow(query, name).Scan(&id)

    if err != nil {
        if err == sql.ErrNoRows {
            // Nie znaleziono producenta o podanej nazwie
            return uuid.Nil, nil
        }
        // Inny błąd
        return uuid.Nil, err
    }

    return id, nil
}