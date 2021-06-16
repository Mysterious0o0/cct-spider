package miit

type Details struct {
	Url string `json:"url"`
}

type GroupData struct {
	Details `json:"data"`
}

type DataResults struct {
	GroupData []GroupData `json:"groupData"`
}

type SearchResult struct {
	DataResults []DataResults `json:"dataResults"`
	Total       int           `json:"total"`
}

type Data struct {
	SearchResult `json:"searchResult"`
}

type JsonMiit struct {
	Data `json:"data"`
}
