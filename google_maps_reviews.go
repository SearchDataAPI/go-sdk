package searchdata

type GoogleMapsReviewsParams struct {
	Engine		string	`url:"engine"`
	ApiKey 		string	`url:"api_key" binding:"required"`
	Q 			string	`url:"q" binding:"required"`
	Location 	string	`url:"location,omitempty"`
	Device 		string	`url:"device,omitempty"`
}

func GoogleMapsReviews(params GoogleMapsReviewsParams) (SearchResult, error) {
	params.Engine = "google_maps_reviews"
	return GetJson(params)
}