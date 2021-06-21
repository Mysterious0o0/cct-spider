package mohurd

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
			UrlSelector: ".table01 a",
		},
	}
	pr.GetPageUrl("href")
}

func GetSecondUrl(url string) {
	s := strings.Split(url, "/")
	fmt.Println(s[3])
	switch s[4] {
	case "xwfbzxft", "zcjd", "gzly", "xf", "cjda", "zfggyfz":
		fmt.Println(url)
	default:
		pr := response.PR{
			Request: request.Request{
				Url:    url,
				Method: http.MethodGet,
			},
			Parse: parse.Parse{
				BaseUrl:     url,
				UrlSelector: "body>table>tbody>tr:nth-child(2)>td>table>tbody>tr:nth-child(1)>td:nth-child(1)>table>tbody>tr>td>a",
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
			PageNumSelector: "body",
		},
	}
	num := pr.GetPageNum("共[0-9]+页")
	if num == 0 {
		num = 40
	}
	for i := 2; i <= num; i++ {
		url := fmt.Sprintf("%s_%v.html", url[:len(url)-len(".html")], i)
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
			UrlSelector: "td[valign='top'][align='middle'] tbody a",
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
			TitleSelector: ".tit",
			TextSelector: ".info .union>p",
			DomainName: "http://www.mohurd.gov.cn",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}