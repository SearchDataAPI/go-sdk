# SearchData Go SDK

SearchData is an API that allows scraping various search engines such as Google, Bing, Yandex, etc. while using rotating proxies to prevent bans. This SDK for Go makes the usage of the API easier to implement in any project you have.

## Installation

You must have Go 1.16 installed.

Run the following command to add the package:

```
go get -u github.com/searchdataapi/go-sdk
```

## Quick Start

Use the following code to test the SDK:

```
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
```