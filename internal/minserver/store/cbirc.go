package store

type DetailsCbirc struct {
	DocId interface{} `json:"docId"`
}

type Rows struct {
	Rows []DetailsCbirc `json:"rows"`
}

type JsonCbirc struct {
	Data Rows `json:"data"`
}

type DataCbirc struct {
	DocTitle string `json:"docTitle"`
	DocClob  string `json:"docClob"`
}

type JsonDetailsCbirc struct {
	DataCbirc `json:"data"`
}

var DetailUrl = "http://www.cbirc.gov.cn/cn/static/data/DocInfo/SelectByDocId/data_docId=%v.json"

