package icApi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	// "log"
	"net/http"
)

var (
    clientKey   string
    clientSecret string
    tokenEndpoint string
)

func init() {
    clientKey = os.Getenv("IC_API_CUSTOMER_KEY")
    clientSecret = os.Getenv("IC_API_CUSTOMER_SECRET")
    tokenEndpoint = os.Getenv("IC_TOKEN_ENDPOINT")
}

func GetToken() (string, error) {
    // Utwórz żądanie HTTP
    request, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBufferString("grant_type=client_credentials"))
    if err != nil {
        return "", err
    }

    // Ustaw nagłówki
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Set("Authorization", basicAuth(clientKey, clientSecret))
	// log.Println(request)
    // Wysłaj żądanie
    response, err := http.DefaultClient.Do(request)
    if err != nil {
        return "", err
    }

   
    // Zamknij ciało odpowiedzi
    defer response.Body.Close()

    // Sprawdź kod odpowiedzi HTTP
    if response.StatusCode < 200 || response.StatusCode >= 300 {
        return "", fmt.Errorf("Nieprawidłowy kod odpowiedzi HTTP: %d", response.StatusCode)
    }

    // Dekoduj odpowiedź JSON
    var token struct {
        Token string `json:"access_token"`
    }
    if err := json.NewDecoder(response.Body).Decode(&token); err != nil {
        return "", err
    }

    // Wyświetl token
    return token.Token, nil
}

func basicAuth(username, password string) string {
    auth := username + ":" + password
    encoded := base64.StdEncoding.EncodeToString([]byte(auth))
    return "Basic " + encoded
}