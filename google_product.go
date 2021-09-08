package searchdata

type GoogleProductParams struct {
	Engine			string	`url:"engine"`
	ApiKey 			string	`url:"api_key" binding:"required"`
	ProductId 		string	`url:"product_id" binding:"required"`
	Location 		string	`url:"location,omitempty"`
	Uule 			string	`url:"uule,omitempty"`
	GoogleDomain 	string	`url:"google_domain,omitempty"`
	Hl 				string	`url:"hl,omitempty"`
	Gl 				string	`url:"gl,omitempty"`
	Start 			int		`url:"start,omitempty"`
	Device 			string	`url:"device,omitempty"`
}

func GoogleProduct(params GoogleProductParams) (SearchResult, error) {
	params.Engine = "google_product"
	return GetJson(params)
}
