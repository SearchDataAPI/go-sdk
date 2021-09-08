package searchdata

type GoogleMapsParams struct {
	Engine			string	`url:"engine"`
	ApiKey 			string	`url:"api_key" binding:"required"`
	Q 				string	`url:"q" binding:"required"`
	Type 			string	`url:"type,omitempty"`
	Data 			string	`url:"data,omitempty"`
	Hl 				string	`url:"hl,omitempty"`
	GoogleDomain 	string	`url:"google_domain,omitempty"`
	Ll 				string	`url:"ll,omitempty"`
}

func GoogleMaps(params GoogleMapsParams) (SearchResult, error) {
	params.Engine = "google_maps"
	return GetJson(params)
}