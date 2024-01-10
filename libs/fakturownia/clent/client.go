package facturownia

import (
	"api/libs/fakturownia/service"

	"net/url"
	"strconv"

)

type ClientSearchParam string 

const (
    ClientName    ClientSearchParam = "name"
    ClientEmail   ClientSearchParam = "email"
    ClientShortcut ClientSearchParam = "shortcut"
    ClientTaxNo   ClientSearchParam = "tax_no"
)

func ListClients(page int, perPage int) ([]byte, error) {
    endpoint := "/clients.json"
    params := "&page=" + strconv.Itoa(page) + "&per_page=" + strconv.Itoa(perPage)
    return facturownia.CallFakturowniaAPI(endpoint, params, "GET", nil)
}

func SearchClients(searchParam ClientSearchParam, searchValue string) ([]byte, error) {
    endpoint := "/clients.json"
    params := string(searchParam) + "=" + url.QueryEscape(searchValue)
    return facturownia.CallFakturowniaAPI(endpoint, params, "GET", nil)
}
