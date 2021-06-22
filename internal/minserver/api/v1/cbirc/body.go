package cbirc

type Details struct {
	DocId interface{} `json:"docId"`
}

type Rows struct {
	Rows []Details `json:"rows"`
}

type JsonCbirc struct {
	Data Rows `json:"data"`
}

type Data struct {
	DocTitle string `json:"docTitle"`
	DocClob  string `json:"docClob"`
}

type JsonDetails struct {
	Data `json:"data"`
}

var FileBaseUrl = "https://www.cbirc.gov.cn/cbircweb/DocInfo/SelectDocByItemIdAndChild?itemId=928&pageSize=20&pageIndex=%v"
var LawsBaseUrl = "https://www.cbirc.gov.cn/cbircweb/DocInfo/SelectDocByItemIdAndChild?itemId=927&pageSize=20&pageIndex=%v"
var DetailUrl = "http://www.cbirc.gov.cn/cn/static/data/DocInfo/SelectByDocId/data_docId=%v.json\n"
