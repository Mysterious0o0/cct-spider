package mod

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
)

func GetFirstUrl(url string) {
	fmt.Println(url) // frist url
	for i := 2; i < 3; i++ {
		url := fmt.Sprintf("%s_%v.htm", url[:len(url)-len(".htm")], i)
		fmt.Println(url) // add other url
	}
}

func GetDetailPageUrl(url string) {
	//baseurl := "http://www.mod.gov.cn/regulatory/"
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: "#main-news-list>li>a",
		},
	}
	pr.GetPageUrl("href")
}

func GetHtmlInfo(url string) (infoMap map[string]string){
	infoMap = make(map[string]string)
	pr := response.PR{
		Request: request.Request{
			Url : url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			DomainName: "http://www.mod.gov.cn/",
			TextSelector: "#article-content>p",
			TitleSelector: ".article-header>h1, .article-header>h3",

		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}
