package searchdata

type YandexSearchParams struct {
	Engine			string	`url:"engine"`
	ApiKey			string	`url:"api_key" binding:"required"`
	Q 				string	`url:"q" binding:"required"`
	Location 		string	`url:"location,omitempty"`
	Language 		string	`url:"language,omitempty"`
	YandexDomain 	string	`url:"yandex_domain,omitempty"`
	Device 			string	`url:"device,omitempty"`
	P 				int		`url:"p,omitempty"`
}

func YandexSearch(options YandexSearchParams) (SearchResult, error)  {
	options.Engine = "yandex"
	return GetJson(options)
}