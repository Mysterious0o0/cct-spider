package mee

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
	"strings"
)

func GetFirstUrl(url string) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: "div[class='outBox zcwj']>div>a",
		},
	}
	pr.GetPageUrl("href")
}

func GetSecondUrl(url string) {
	s := strings.Split(url, "/")
	fmt.Println(s[4])
	switch s[4] {
	case "zyygwj", "gwywj", "xzspwj":
		fmt.Println(url)
	default:
		pr := response.PR{
			Request: request.Request{
				Url:    url,
				Method: http.MethodGet,
			},
			Parse: parse.Parse{
				BaseUrl:     url,
				UrlSelector: "span[class='mobile_none']>a",
			},
		}
		pr.GetPageUrl("href")
	}
}

func GetThreeUrl(url string) {
	fmt.Println(url) // frist url
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: ".slideTxtBoxgsf script",
		},
	}
	num := pr.GetPageNum("var countPage = \\d+//")
	if num == 0 {
		num = 40
	}
	for i := 1; i < num; i++ {
		url := fmt.Sprintf("%sindex_%v.shtml", url, i)
		fmt.Println(url) // other url
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
			UrlSelector: "#div>li>a",
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
			TitleSelector: "h1, h2",
			TextSelector: ".Custom_UnionStyle p, .Custom_UnionStyle div, .content_body_box>p, .content_body_box>div, .neiright_JPZ_GK_CP>p",
			DomainName: "http://www.mee.gov.cn",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}

