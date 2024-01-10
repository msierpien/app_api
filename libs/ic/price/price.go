package icApi

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)



func ImportPriceCSV() {
	fileDir := "./assets/prices/"
	csvPattern := regexp.MustCompile(`Wholesale_Pricing_(\d{4}-\d{2}-\d{2})\.csv`)

	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		log.Fatalf("Błąd podczas czytania katalogu %s: %v", fileDir, err)
	}

	var latestFile os.FileInfo
	var latestDate time.Time

	for _, file := range files {
		matches := csvPattern.FindStringSubmatch(file.Name())
		if len(matches) == 2 {
			dateStr := matches[1]
			date, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				log.Printf("Błąd podczas parsowania daty z nazwy pliku %s: %v", file.Name(), err)
				continue
			}

			if date.After(latestDate) {
				latestDate = date
				latestFile = file
			}
		}
	}

	if latestFile != nil {
		log.Printf("Najnowszy plik CSV to: %s", latestFile.Name())
		// Tutaj możesz wczytać i przetworzyć plik
	} else {
		log.Println("Nie znaleziono plików CSV pasujących do wzorca.")
	}

}