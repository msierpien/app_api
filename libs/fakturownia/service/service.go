package facturownia

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"


)



func CallFakturowniaAPI(endpoint string, params string, method string, payload interface{}) ([]byte, error) {
	apiKey := os.Getenv("FAKTUROWNIA_TOKEN_API")
	key := "?api_token=" + apiKey
    // apiUrl := os.Getenv("FAKTUROWNIA_BASE_URL") + endpoint + key
    var apiUrl string
    switch method {
    case "GET":
        apiUrl = os.Getenv("FAKTUROWNIA_BASE_URL") + endpoint + key + params
    case "POST":
        apiUrl = os.Getenv("FAKTUROWNIA_BASE_URL") + endpoint
    }

    jsonValue, _ := json.Marshal(payload)
    request, err := http.NewRequest(method, apiUrl,  bytes.NewBuffer(jsonValue))
	log.Println(request)
    if err != nil {
        return nil, err
    }

    request.Header.Set("Content-Type", "application/json")


    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}