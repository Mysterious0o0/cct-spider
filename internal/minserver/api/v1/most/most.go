package most

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/respont"
	"net/http"
	"strings"
)

func GetFirstUrl(url string) {
	baseUrl := "http://www.most.gov.cn/kjzc/gjkjzc/"
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     baseUrl,
			UrlSelector: ".list>a",
			Suffix:      "index.html",
		},
	}
	pr.GetPageUrl("href")
}


func GetPageUrlList(url string) {
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: "script[language='JavaScript']",
		},
	}
	num := pr.GetPageNum("var countPage = \\d+")
	fmt.Println(num)
	if num == 0 {
		num = 20
	}
	fmt.Println(url) // first page url
	for i := 1; i < num; i++ {
		otherUrl := fmt.Sprintf("%s_%v.html", url[:len(url)-len(".html")], i)
		fmt.Println(otherUrl) // other page url
	}
}

func GetDetailPageUrl(url string) {
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: ".list>ul a",
		},
	}
	pr.GetPageUrl("href")
}


func GetHtmlInfo(url string) (infoMap map[string]string) {
	infoMap = make(map[string]string)
	var pr respont.PR
	r := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	s := strings.Split(url, "/")
	fmt.Println(s[3])
	switch s[3] {
	case "xxgk":
		pr = respont.PR{
			Request: r,
			Parse: parse.Parse{
				DomainName:    "http://www.most.gov.cn/",
				TextSelector:  "#Zoom p",
				TitleSelector: ".xxgk_title",
			},
		}
	default:
		pr = respont.PR{
			Request: r,
			Parse: parse.Parse{
				DomainName:    "http://www.most.gov.cn/",
				TextSelector:  "#Zoom p",
				TitleSelector: "#Title",
			},
		}
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}



