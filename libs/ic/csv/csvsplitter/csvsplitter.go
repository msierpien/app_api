package csvsplitter

import (
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
)

func SplitCSV(records [][]string, rowsPerFile int, outputDir string) error {
    if len(records) == 0 {
        return fmt.Errorf("brak rekordów do przetworzenia")
    }

    // Zakładamy, że pierwszy wiersz to nagłówki
    headers := records[0]

    fileCount := 0
    for i := 1; i < len(records); i += rowsPerFile { // Rozpoczynamy od 1, aby pominąć nagłówki
        end := i + rowsPerFile
        if end > len(records) {
            end = len(records)
        }

        fileName := filepath.Join(outputDir, fmt.Sprintf("part_%d.csv", fileCount))
        file, err := os.Create(fileName)
        if err != nil {
            return err
        }

        csvWriter := csv.NewWriter(file)
        csvWriter.Comma = ';' // Ustawienie separatora na średnik

        // Zapisz nagłówki w każdym nowym pliku
        if err := csvWriter.Write(headers); err != nil {
            file.Close()
            return err
        }

        for _, record := range records[i:end] {
            if err := csvWriter.Write(record); err != nil {
                file.Close()
                return err
            }
        }

        csvWriter.Flush()
        file.Close()

        if err := csvWriter.Error(); err != nil {
            return err
        }

        fileCount++
    }

    return nil
}