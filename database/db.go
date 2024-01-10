package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

// Ładuje zmienne środowiskowe z pliku .env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Błąd wczytywania pliku .env")
	}
}

// Connect to the database and return the connection
func Connect() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Błąd podczas otwierania bazy danych: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Błąd podczas pingowania bazy danych: %v\n", err)
	}

	fmt.Println("Pomyślnie połączono z bazą danych")
	return db
}
