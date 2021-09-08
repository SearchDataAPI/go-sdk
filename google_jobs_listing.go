package searchdata

type GoogleJobsListingParams struct {
	Engine	string	`url:"engine"`
	ApiKey 	string	`url:"api_key" binding:"required"`
	Q 		string	`url:"q" binding:"required"`
}

func GoogleJobsListing(params GoogleJobsListingParams) (SearchResult, error) {
	params.Engine = "google_jobs_listing"
	return GetJson(params)
}
