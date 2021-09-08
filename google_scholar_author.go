package searchdata

type GoogleScholarAuthorParams struct {
	Engine		string	`url:"engine"`
	ApiKey 		string	`url:"api_key" binding:"required"`
	AuthorId 	string	`url:"author_id" binding:"required"`
	Sort 		string	`url:"sort,omitempty"`
	ViewOp 		string	`url:"view_op,omitempty"`
	CitationId 	string	`url:"citation_id,omitempty"`
	Offset 		int		`url:"offset,omitempty"`
	Num 		int		`url:"num,omitempty"`
	Hl 			string	`url:"hl,omitempty"`
	Device 		string	`url:"device,omitempty"`
}

func GoogleScholarAuthor(params GoogleScholarAuthorParams) (SearchResult, error) {
	params.Engine = "google_scholar_author"
	return GetJson(params)
}
