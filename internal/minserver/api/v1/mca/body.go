package mca

type Details struct {
	Title       string `json:"title"`
	HtmlContent string `json:"htmlContent"`
}

type JsonMca struct {
	ResultMap    []Details `json:"resultMap"`
}

