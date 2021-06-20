package mca

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"net/http"
	"strings"
)

func GetHtmlInfo(url string) (info map[string]string) {
	info = make(map[string]string)
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	s := string(b)
	var delChar = []string{"\\n", "\n", "\t", "\\t", "\\r", "\r"}
	for _, v := range delChar{
		s = strings.Replace(s, v, "", -1)
	}
	if err != nil {
		return
	}
	var j JsonMca
	err = json.Unmarshal([]byte(s[1:len(s)-2]), &j)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range j.ResultMap{
		info[v.Title] = strings.Join(_getInfo(v.HtmlContent), "")
	}
	//fmt.Println(info)
	return info
}

func _getInfo(html string) (data []string){
	p := parse.Parse{
		Html:            html,
		TextSelector:    "p",
	}
	_, data = p.GetTextByParseHtml()
	return
}

