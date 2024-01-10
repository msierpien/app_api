package facturownia

import (
	"api/graph/model"
	"api/libs/fakturownia/service"
	"encoding/json"
	"log"
	"os"
	"strconv"
)



type Payload struct {
    APIToken string   `json:"api_token"`
    Invoice  model.InvoiceInput  `json:"invoice"`
}

func InvoiceClient(client_id string)( []byte, error) {

	endpoints := "/invoices.json"
	params :=  "&client_id=" + client_id

	doc, err := facturownia.CallFakturowniaAPI(endpoints, params , "GET", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return doc, err
}
func InvoiceID(id int)( []byte, error) {

	endpoints := "/invoices/"+strconv.Itoa(id)+".json"
	params :=  ""

	doc, err := facturownia.CallFakturowniaAPI(endpoints, params , "GET", nil)
	if err != nil {
		log.Println(err)
		return doc, err
	}
	return doc, err
}
func InvoiceAdd( invoice model.InvoiceInput) ( []byte, error)  {
	endpoint := "/invoices.json"
	apiToken := os.Getenv("FAKTUROWNIA_TOKEN_API")
    payload := Payload{
        APIToken: apiToken,
        Invoice:  invoice,
    }

    jsonPayload, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error marshalling payload: %v\n", err)
        return nil, err
    }
    log.Printf("Sending payload: %s\n", jsonPayload)

    doc, err := facturownia.CallFakturowniaAPI(endpoint, "", "POST", &payload)
    if err != nil {
        log.Println("Error calling Fakturownia API:", err)
        return nil, err
    }
    log.Println("Received response:", string(doc))
    return doc, nil
}