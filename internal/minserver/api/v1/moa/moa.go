package moa

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
)

func GetPageUrlList(url string) {
	fmt.Println(url) // frist url
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: ".next script",
		},
	}
	num := pr.GetPageNum("var countPage = [0-9]+//")
	if num == 0 {
		num = 10
	}
	for i := 1; i < num; i++ {
		url := fmt.Sprintf("%sindex_%v.htm", url, i)
		fmt.Println(url) // add other url
	}
}

func GetDetailPageUrl(url string) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: ".commonlist>li>a",
		},
	}
	pr.GetPageUrl("href")
}

func GetHtmlInfo(url string) (infoMap map[string]string) {
	infoMap = make(map[string]string)
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			TitleSelector: "h2",
			TextSelector:  ".gsj_htmlcon_bot>p, .Custom_UnionStyle>p",
			DomainName:    "http://www.moa.gov.cn",
		},
	}
	infoMap = pr.GetHtmlInfo()
	//if len(infoMap) != 0{
	//	fmt.Println("存入")
	//}
	fmt.Println(infoMap)
	return
}

