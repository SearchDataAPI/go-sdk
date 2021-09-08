package searchdata

type GoogleReverseImageParams struct {
	Engine			string	`url:"engine"`
	ApiKey 			string	`url:"api_key" binding:"required"`
	ImageUrl 		string	`url:"image_url" binding:"required"`
	Location 		string	`url:"location,omitempty"`
	Uule 			string	`url:"uule,omitempty"`
	GoogleDomain 	string	`url:"google_domain,omitempty"`
	Hl 				string	`url:"hl,omitempty"`
	Gl 				string	`url:"gl,omitempty"`
	Lr 				string	`url:"lr,omitempty"`
	Device 			string	`url:"device,omitempty"`
}

func GoogleReverseImage(params GoogleReverseImageParams) (SearchResult, error) {
	params.Engine = "google_reverse_image"
	return GetJson(params)
}
