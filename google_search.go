package searchdata

type GoogleSearchParams struct {
	Engine			string	`url:"engine"`
	ApiKey 			string	`url:"api_key" binding:"required"`
	Q 				string	`url:"q" binding:"required"`
	Location 		string	`url:"location,omitempty"`
	Ludocid 		string	`url:"ludocid,omitempty"`
	Lr 				string	`url:"lr,omitempty"`
	Hl 				string	`url:"hl,omitempty"`
	Gl 				string	`url:"gl,omitempty"`
	Uule 			string	`url:"uule,omitempty"`
	GoogleDomain 	string	`url:"google_domain,omitempty"`
	Start 			string	`url:"start,omitempty"`
	Num 			int		`url:"num,omitempty"`
	Tbm 			string	`url:"tbm,omitempty"`
	SortBy 			string	`url:"sort_by,omitempty"`
	Cookie 			string	`url:"cookie,omitempty"`
	Lsig			string	`url:"lsig,omitempty"`
	TimePeriod 		string	`url:"time_period,omitempty"`
	Tbs 			string	`url:"tbs,omitempty"`
	FlattenResults 	int		`url:"flatten_results,omitempty"`
	Safe 			int		`url:"safe,omitempty"`
	Filter 			int		`url:"filter,omitempty"`
	Nfpr 			int		`url:"nfpr,omitempty"`
	Device 			string	`url:"device,omitempty"`
	EmptyResults 	int		`url:"empty_results,omitempty"`
}

func GoogleSearch(options GoogleSearchParams) (SearchResult, error)  {
	options.Engine = "google"
	return GetJson(options)
}