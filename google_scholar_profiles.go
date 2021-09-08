package searchdata

type GoogleScholarProfilesParams struct {
	Engine			string	`url:"engine"`
	ApiKey 			string	`url:"api_key" binding:"required"`
	Q 				string	`url:"q" binding:"required"`
	AfterAuthor 	string	`url:"after_author,omitempty"`
	BeforeAuthor 	string	`url:"before_author,omitempty"`
	Hl 				string	`url:"hl,omitempty"`
}

func GoogleScholarProfiles(params GoogleScholarProfilesParams) (SearchResult, error) {
	params.Engine = "google_scholar_profiles"
	return GetJson(params)
}
