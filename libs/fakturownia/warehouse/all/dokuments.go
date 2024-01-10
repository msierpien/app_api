package facturownia

import (
	"api/libs/fakturownia/service"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Błąd wczytywania pliku .env")
	}
}




func AllWarehouse(page int, perPage int)  {
	endpoint := "/warehouse_documents.json?"
    params := "page=" + strconv.Itoa(page) + "&per_page=" + strconv.Itoa(perPage)

	dock, err := facturownia.CallFakturowniaAPI( endpoint, params, "GET", nil) 
	{
		if err != nil {
			log.Println(err)
			return
		}
	}
	log.Println(string(dock))

}
func AddWarehouse( payload interface{})( []byte, error)  {
	endpoint := "/warehouse_documents.json"


	dock, err := facturownia.CallFakturowniaAPI(endpoint, "", "POST", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(string(dock))
	return dock, err
}

func EditWarehouse( payload interface{}, idDoc int8)( []byte, error)  {
	endpoint := "/warehouse_documents"+string(rune(idDoc))+".json"
	doc, err := facturownia.CallFakturowniaAPI(endpoint, "", "PUT", payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(string(doc))
	return doc, err
}

func DeleteWarehouse( idDoc int8)( []byte, error)  {
	endpoint := "/warehouse_documents"+string(rune(idDoc))+".json"
	doc, err := facturownia.CallFakturowniaAPI(endpoint, "", "DELETE", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(string(doc))
	return doc, err
}