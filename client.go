package searchdata

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

type SearchResult map[string]interface{}

func request(o interface{}) (*http.Response, error) {

	values, _ := query.Values(o)
	queryParams := values.Encode()

	resp, err := http.Get("https://api.searchdata.io/v1?" + queryParams)

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func decodeJSON(body []byte) (SearchResult, error) {

	response := SearchResult{}

	err := json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return response, nil

}

func GetJson(o interface{}) (SearchResult, error) {

	resp, err := request(o)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	decoded, err := decodeJSON(body)

	if err != nil {
		return nil, err
	}

	return decoded, nil

}