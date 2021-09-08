package searchdata

type BingSearchParams struct {
	Engine		string	`url:"engine"`
	ApiKey 		string	`url:"api_key" binding:"required"`
	Q 			string 	`url:"q" binding:"required"`
	Location 	string 	`url:"location,omitempty"`
	Cc 			string 	`url:"cc,omitempty"`
	Offset 		int		`url:"offset,omitempty"`
	Count 		int		`url:"count,omitempty"`
	Device 		int		`url:"device,omitempty"`
	SafeSearch 	string	`url:"safe_search,omitempty"`
	SetLang 	string	`url:"set_lang,omitempty"`
	Mkt 		string	`url:"mkt,omitempty"`
}

func BingSearch(params BingSearchParams) (SearchResult, error) {
	params.Engine = "bing"
	return GetJson(params)
}