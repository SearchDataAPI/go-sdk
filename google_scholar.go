package searchdata

type GoogleScholarParams struct {
	Engine		string	`url:"engine"`
	ApiKey 		string	`url:"api_key" binding:"required"`
	Q 			string	`url:"q" binding:"required"`
	Cites 		string	`url:"cites,omitempty"`
	Cluster 	string	`url:"cluster,omitempty"`
	AsYlo 		int		`url:"as_ylo,omitempty"`
	AsYhi 		int		`url:"as_yhi,omitempty"`
	Scisbd 		int		`url:"scisbd,omitempty"`
	AsVis 		int		`url:"as_vis,omitempty"`
	Lr 			string	`url:"lr,omitempty"`
	Hl 			string	`url:"hl,omitempty"`
	Safe 		int		`url:"safe,omitempty"`
	Start 		int		`url:"start,omitempty"`
	Num 		int		`url:"num,omitempty"`
}

func GoogleScholar(params GoogleScholarParams) (SearchResult, error) {
	params.Engine = "google_scholar"
	return GetJson(params)
}

