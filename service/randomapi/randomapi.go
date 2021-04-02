package randomapi

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetItemFunc of function type
type GetItemFunc func(string) []byte

// RandomAPI has the common connection part
type RandomAPI struct {
	GetItem GetItemFunc
}

// GetItemFunc ss
func (c RandomAPI) GetItemFunc(resource string) []byte {
	resp, err := http.Get("https://random-data-api.com/api/" + resource)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
