package searchdata

type GoogleAutocompleteParams struct {
	Engine	string	`url:"engine"`
	ApiKey 	string	`url:"api_key" binding:"required"`
	Q 		string	`url:"q" binding:"required"`
}

func GoogleAutocomplete(params GoogleAutocompleteParams) (SearchResult, error) {
	params.Engine = "google_autocomplete"
	return GetJson(params)
}