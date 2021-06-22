package mwr

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
			PageNumSelector: ".fy script",
		},
	}
	num := pr.GetPageNum("var countPage = [0-9]+//")
	if num == 0 {
		num = 100
	}
	for i := 1; i < num; i++ {
		url := fmt.Sprintf("%sindex_%v.html", url, i)
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
			UrlSelector: ".slnewsconlist>li>a",
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
			TitleSelector: "h1",
			TextSelector:  ".TRS_Editor>p",
			DomainName:    "http://www.mwr.gov.cn",
			Suffix:        ".html",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}
