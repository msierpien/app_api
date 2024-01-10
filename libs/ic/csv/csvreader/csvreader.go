package csvreader

import (
    "encoding/csv"
    "io"
    "os"
)

func ReadCSV(filePath string) ([][]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.Comma = ';' // Ustawienie separatora na średnik
    var records [][]string

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
                continue // Pomija wiersz z błędem liczby pól
            }
            return nil, err
        }
        records = append(records, record)
    }

    return records, nil
}