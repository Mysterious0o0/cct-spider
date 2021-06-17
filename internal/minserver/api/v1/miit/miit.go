package miit

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/respont"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/urlprocess"
	"net/http"
)

func GetFirstUrl(url string) {
	for i := 0; i <= 100; i++ {
		url := fmt.Sprintf("%s%v", url, i)
		fmt.Println(url)
	}
}

func GetDetailPageUrl(url string, baseUrl string) {
	baseUrl = "https://www.miit.gov.cn"
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	if err != nil {
		return
	}
	var j JsonMiit
	err = json.Unmarshal(b, &j)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(j.Data.Total) //总文件数
	for _, v := range j.Data.DataResults {
		fmt.Println(urlprocess.UrlJoint(baseUrl, v.GroupData[0].Url))
		//fmt.Println(baseUrl + v.GroupData[0].Url)
	}
}

func GetHtmlInfo(url string, baseUrl string) (infoMap map[string]string) {
	infoMap = make(map[string]string)
	//var info []string
	//req := request.Request{
	//	Url : urlprocess,
	//	Method: http.MethodGet,
	//}
	//html, err := req.Visit()
	//if err != nil {
	//	return
	//}
	////fmt.Println(string(html))
	//p := parse.Parse{
	//	Html:     string(html),
	//	UrlSelector: "iframe",
	//	TitleSelector: "#con_title",
	//	TextSelector: "#con_con>p",
	//}
	//title, data := p.GetTextByParseHtml()
	//info = append(info, data...)
	//fileurl, b := p.GetOneUrlByParseHtml("fileurl")
	//if b {
	//	fmt.Println(baseUrl + fileurl)
	//	getPDFInfo(baseUrl + fileurl)
	//}
	//if _, ok := infoMap[title]; !ok{
	//	infoMap[title] = strings.Join(info, "")
	//}
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			UrlSelector:   "iframe",
			TitleSelector: "#con_title",
			TextSelector:  "#con_con>p",
			DomainName:    "https://www.miit.gov.cn/",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}

func getPDFInfo(url string) (info []string) {
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	html, err := req.Visit()
	if err != nil {
		return
	}
	fmt.Println(string(html))
	return
}
