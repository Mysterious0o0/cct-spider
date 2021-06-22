package cbirc

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"net/http"
	"strings"
)

func GetPageUrlList(url string) {

	for i := 1; i <= 1000; i++ {
		req := request.Request{
			Url:    fmt.Sprintf(url, i),
			Method: http.MethodGet,
		}
		b, err := req.Visit()
		if err != nil {
			break
		}
		var j JsonCbirc
		err = json.Unmarshal(b, &j)
		if err != nil {
			fmt.Println(err)
		}
		if len(j.Data.Rows) == 0 {
			break
		}
		for _, v := range j.Data.Rows {
			fmt.Printf(DetailUrl, v.DocId)
		}
	}
}


func GetHtmlInfo(url string) (infoMap map[string]string) {
	infoMap = make(map[string]string)
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	if err != nil {
		return
	}
	var j JsonDetails
	err = json.Unmarshal(b, &j)
	if err != nil {
		fmt.Println(err)
	}

	p := parse.Parse{
		Html: j.DocClob,
		TextSelector:  "p",
	}
	_, info := p.GetTextByParseHtml()
	infoMap[j.DocTitle] = strings.Join(info, "")
	fmt.Println(infoMap)
	return
}