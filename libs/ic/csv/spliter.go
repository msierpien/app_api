package spliter

import (
    "api/libs/ic/csv/csvreader"   // Zaimportuj swój pakiet csvreader
    "api/libs/ic/csv/csvsplitter" // Zaimportuj swój pakiet csvsplitter
    "log"
)

func Spliter() {
    filePath := "./assets/data//product.csv"
    outputDir := "./assets/split"

    log.Println("Czytanie pliku CSV...")
    records, err := csvreader.ReadCSV(filePath)
    if err != nil {
        log.Fatalf("Błąd podczas czytania pliku CSV: %v", err)
    }

    log.Printf("Znaleziono %d wierszy do podziału.\n", len(records))

    log.Println("Dzielenie pliku CSV...")
    err = csvsplitter.SplitCSV(records, 20000, outputDir)
    if err != nil {
        log.Fatalf("Błąd podczas dzielenia pliku CSV: %v", err)
    }

    log.Println("Plik CSV został pomyślnie podzielony.")
}
