package searchdata

type GoogleScholarCiteParams struct {
	Engine	string	`url:"engine"`
	ApiKey 	string	`url:"api_key" binding:"required"`
	Q 		string	`url:"q" binding:"required"`
	Device 	string	`url:"device,omitempty"`
}

func GoogleScholarCite(params GoogleScholarCiteParams) (SearchResult, error) {
	params.Engine = "google_scholar_cite"
	return GetJson(params)
}
