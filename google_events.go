package searchdata

type GoogleEventsParams struct {
	Engine			string	`url:"engine"`
	ApiKey 			string	`url:"api_key" binding:"required"`
	Q 				string	`url:"q" binding:"required"`
	Location 		string	`url:"location,omitempty"`
	Device 			string	`url:"device,omitempty"`
	Uule 			string	`url:"uule,omitempty"`
	GoogleDomain 	string	`url:"google_domain,omitempty"`
	Hl 				string	`url:"hl,omitempty"`
	Gl 				string	`url:"gl,omitempty"`
}

func GoogleEvents(params GoogleEventsParams) (SearchResult, error) {
	params.Engine = "google_events"
	return GetJson(params)
}