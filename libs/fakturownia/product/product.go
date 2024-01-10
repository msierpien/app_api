package fakturownia

import (
	"api/libs/fakturownia/service"
)
func ProductAll() ( []byte, error)  {
	endpoint := "/products.json"
	params :=  ""
	doc, err := facturownia.CallFakturowniaAPI(endpoint, params, "GET", nil)
	if err != nil {
		return nil, err
	}
	return doc, err
}