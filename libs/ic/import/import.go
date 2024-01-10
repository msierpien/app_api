package importCSV

import (
	"api/database"
	"api/database/models/brands"
	"api/database/models/car_parts"
	"api/database/models/product_suppliers"
	"api/libs/ic/csv/csvreader"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)


type ColumnProcessor func(*modelCarParts.CarPart, string, *sql.DB)

func ImportCSV() {
    const fileDir = "./assets/split/"
	database := db.Connect()
	defer database.Close()


    // Wylistuj wszystkie pliki w katalogu
    files, err := os.ReadDir(fileDir)
    if err != nil {
        log.Fatalf("Błąd podczas listowania plików: %v", err)
    }

    // Przeczytaj każdy plik CSV
    for _, file := range files {
        if filepath.Ext(file.Name()) == ".csv" {
            filePath := filepath.Join(fileDir, file.Name())
            records, err := csvreader.ReadCSV(filePath)
            if err != nil {
				log.Printf("Błąd podczas czytania pliku %s: %v", filePath, err)
                continue // Przechodzi do następnego pliku
            }
			headers := records[0]
			processRecords(headers, records[1:], database)


			//  pierwszy plik CSV jeśli jest break
			break
			 
		// Usuń plik po przetworzeniu
			err = os.Remove(filePath)
			if err != nil {
				log.Printf("Błąd podczas usuwania pliku %s: %v", filePath, err)
			} else {
				log.Printf("Plik %s został usunięty", filePath)
			}
        }
    }
}




var columnProcessors = map[string]ColumnProcessor{
    "TOW_KOD":             processTowKod,
    "IC_INDEX":            processIcIndex,
    "TEC_DOC":             processTecDoc,
    "TEC_DOC_PROD":        processTecDocProd,
    "ARTICLE_NUMBER":      processArticleNumber,
    "MANUFACTURER":        processManufacturer,
    "SHORT_DESCRIPTION":   processShortDescription,
    "DESCRIPTION":         processDescription,
    "BARCODES":            processBarcodes,
    "PACKAGE_WEIGHT":      processPackageWeight,
    "PACKAGE_LENGTH":      processPackageLength,
    "PACKAGE_WIDTH":       processPackageWidth,
    "PACKAGE_HEIGHT":      processPackageHeight,
    "CUSTOM_CODE":         processCustomCode,
    // ... inne mapowania, jeśli są potrzebne
}

// func processTowKod(value string) {
//     // Logika przetwarzania dla TOW_KOD
// 	log.Printf("TOW_KOD: %s", value)
// }

func processTowKod(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Przykładowe zapytanie SQL
    // _, err := db.Exec("INSERT INTO twoja_tabela (kolumna) VALUES ($1)", value)
    // if err != nil {
    //     log.Printf("Błąd podczas wstawiania do bazy danych: %v", err)
    // }
	carPart.TowKod = value
}

func processIcIndex(carPart *modelCarParts.CarPart, value string, db *sql.DB ) {
    // Logika przetwarzania dla IC_INDEX
	// log.Printf("IC_INDEX: %s", value)
	carPart.IcIndex = value
}

func processTecDoc(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla TEC_DOC
	// log.Printf("TEC_DOC: %s", value)
	carPart.TecDocArt = value
}

func processTecDocProd(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla TEC_DOC_PROD
	carPart.TecDocBrand = value
}

func processArticleNumber(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla ARTICLE_NUMBER
	carPart.SKU = value
}

func processManufacturer(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
	brandId, err := modelBrands.GetBrandIDByName(db, value)
	if err != nil {
		log.Printf("Błąd podczas pobierania ID producenta: %v", err)
	
	}

	if brandId == uuid.Nil {
		brand := modelBrands.Brand{
			ID: uuid.New(),
			Name: value,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			}
			newBrandId, err := modelBrands.InsertBrand(db, &brand)

			if err != nil {
				log.Printf("Błąd podczas wstawiania producenta: %v", err)
			}
			carPart.BrandId = newBrandId

			// log.Printf("Producent %s został dodany", brand.Name)
	} else {
		carPart.BrandId = brandId
		// log.Printf("Producent %s już istnieje", value)
	}
	
}

func processShortDescription(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla SHORT_DESCRIPTION
	jsonStructure := map[string]map[string]string{
        "pl": {
            "name": value,
        },
    }
	jsonData, err := json.Marshal(jsonStructure)
    if err != nil {
        log.Fatalf("Błąd podczas serializacji do JSON: %v", err)
    }

	carPart.ShortDesc = json.RawMessage(jsonData)
}

func processDescription(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla DESCRIPTION
	jsonStructure := map[string]map[string]string{
        "pl": {
            "name": value,
        },
    }
	jsonData, err := json.Marshal(jsonStructure)
    if err != nil {
        log.Fatalf("Błąd podczas serializacji do JSON: %v", err)
    }

	carPart.Description = json.RawMessage(jsonData)
}

func processBarcodes(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    array := strings.Split(value, ",")
	carPart.EAN = array


}

func processPackageWeight(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
	// Logika przetwarzania dla PACKAGE_WEIGHT
	value = strings.Replace(value, ",", ".", -1)
	if value == "" {
        value = "0"
    }
    num, err := strconv.ParseFloat(value, 64) // 64 oznacza, że używamy float64
    if err != nil {
        log.Fatalf("Błąd podczas konwersji na float64 dla SKU %s: %v", carPart.SKU, err)
    }
    carPart.Weight = num
}

func processPackageLength(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla PACKAGE_LENGTH
	value = strings.Replace(value, ",", ".", -1)
	if value == "" {
        value = "0"
    }
    num, err := strconv.ParseFloat(value, 64) // 64 oznacza, że używamy float64
    if err != nil {
        log.Fatalf("Błąd podczas konwersji na float64 dla SKU %s: %v", carPart.SKU, err)
    }
    carPart.Length = num

}


func processPackageWidth(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla PACKAGE_WIDTH
	value = strings.Replace(value, ",", ".", -1)
	if value == "" {
        value = "0"
    }
    num, err := strconv.ParseFloat(value, 64) // 64 oznacza, że używamy float64
    if err != nil {
        log.Fatalf("Błąd podczas konwersji na float64 dla SKU %s: %v", carPart.SKU, err)
    }
    carPart.Width = num
}

func processPackageHeight(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla PACKAGE_HEIGHT
	value = strings.Replace(value, ",", ".", -1)
	if value == "" {
        value = "0"
    }
    num, err := strconv.ParseFloat(value, 64) // 64 oznacza, że używamy float64
    if err != nil {
        log.Fatalf("Błąd podczas konwersji na float64 dla SKU %s: %v", carPart.SKU, err)
    }
    carPart.Height = num
}

func processCustomCode(carPart *modelCarParts.CarPart, value string, db *sql.DB) {
    // Logika przetwarzania dla CUSTOM_CODE
	value = strings.Replace(value, ",", ".", -1)
	if value == "" {
		value = "0"
	}
	num, err := strconv.Atoi(value) 
	if err != nil {
		log.Fatalf("Błąd podczas konwersji na int dla SKU %s: %v", carPart.SKU, err)
	}
	carPart.CodeCE = num

}


func processRecords(headers []string, records [][]string, db *sql.DB) {
	for _, record := range records {
		
		carPart := modelCarParts.CarPart{
			ID: uuid.New(),
			Visibility: false,
			Status: "New Product",
			Side: json.RawMessage("{}"), // Pusty obiekt JSON
			Slug: json.RawMessage("{}"), // Pusty obiekt JSON
			Name: json.RawMessage("{}"), // Pusty obiekt JSON
			Type: "CARPARTS",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			}
			// Tworzenie struktury danych dla JSON
			historyData := map[string]map[string]string{
				"version": {
					"name": "V 1.0",
				},
			}

			
			
		

			
			// Konwersja sideData na JSON
			historyJSON, err := json.Marshal(historyData)
			if err != nil {
				log.Fatalf("Błąd podczas serializacji History do JSON: %v", err)
			}
			
			carPart.History = json.RawMessage(historyJSON)
			
			
			for i, value := range record {
			

				columnName := headers[i]
				if processor, ok := columnProcessors[columnName]; ok {
					processor(&carPart, value, db)
					} else {
						// log.Printf("Nieznana kolumna: %s", columnName)
					}
				}
				exist, err := modelCarParts.ExistsByTowKod(db, carPart.TowKod)
				if err != nil {
					log.Fatalf("Błąd podczas sprawdzania istnienia CarPart: %v", err)
				}
				if exist {
					// log.Printf("CarPart %s już istnieje", carPart.TowKod)
					} else {
						carPartId, err := modelCarParts.InsertCarPart(db, &carPart)
						if err != nil {
							log.Fatalf("Błąd podczas wstawiania CarPart: %v", err)
						}
					
					
						carPartSuppliers := models.ProductSuppliers{
						
							CarPartId:  carPartId, 
							SupplierId: uuid.MustParse("4c5db981-6451-4daf-a023-f11cd2945efa"), // stała wartość UUID dla SupplierId
							IndexSuppliers: carPart.TowKod,
							CreatedAt:  time.Now(),
							UpdatedAt:  time.Now(),
						}
						
				models.InsertCarPartSuppliers(db, &carPartSuppliers)
				// log.Printf("CarPart %s został dodany", carPart.SKU)
				
			}

    }
}


func Import(){
	ImportCSV()
}