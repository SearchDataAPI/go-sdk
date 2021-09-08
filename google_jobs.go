package searchdata

type GoogleJobsParams struct {
	Engine			string	`url:"engine"`
	ApiKey 			string	`url:"api_key" binding:"required"`
	Q 				string	`url:"q" binding:"required"`
	Location 		string	`url:"location,omitempty"`
	Uule 			string	`url:"uule,omitempty"`
	GoogleDomain 	string	`url:"google_domain,omitempty"`
	Hl 				string	`url:"hl,omitempty"`
	Gl 				string	`url:"gl,omitempty"`
	Start 			int		`url:"start,omitempty"`
	Device 			string	`url:"device,omitempty"`
}

func GoogleJobs(params GoogleJobsParams) (SearchResult, error) {
	params.Engine = "google_jobs"
	return GetJson(params)
}