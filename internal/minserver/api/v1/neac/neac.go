package neac

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
)

func GetFirstUrl(url string) {
	fmt.Println(url) // frist url
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: ".w1 script",
		},
	}
	num := pr.GetPageNum(",\\d+,")
	if num == 0 {
		num = 100
	}
	for i := 2; i <= num; i++ {
		url := fmt.Sprintf("%s_%v.shtml", url[:len(url)-len(".shtml")], i)
		fmt.Println(url) // add other url
	}
}

func GetDetailPageUrl(url string) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			BaseUrl: url,
			UrlSelector: ".w1 span>a",
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
			TitleSelector: ".p1",
			TextSelector: "div[class='wzy p3']>p",
			DomainName: "https://www.neac.gov.cn",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}