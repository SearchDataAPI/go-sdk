package main

import (
	"fmt"
	"github.com/searchdataapi/go-sdk"
)

func main()  {

	response, err := searchdata.GoogleSearch(searchdata.GoogleSearchParams{
		ApiKey: "{API_KEY}",
		Q: "pizza",
	})

	if err != nil {
		panic(err)
	}

	results := response["organic_results"].([]interface{})
	firstResult := results[0].(map[string]interface{})
	fmt.Println(firstResult["title"].(string))

}

