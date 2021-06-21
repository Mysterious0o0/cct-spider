package moe

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
)

func GetPageUrlList(url string) {
	num := 20
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: "head>script",
		},
	}
	// "var recordCount = [0-9]+", " var pageSize = [0-9]+"
	countR := "var recordCount = [0-9]+"
	sizeR := "var pageSize = [0-9]+"
	count, size := pr.GetCountAndSize(countR, sizeR)
	fmt.Println(count, size)
	if count != 0 && size != 0 {
		num = count / size + 1
	}
	fmt.Println(url) // first page url
	for i := 1; i < num; i++ {
		otherUrl := fmt.Sprintf("%sindex_%v.html", url, i)
		fmt.Println(otherUrl) // other page url
	}
}

func GetDetailPageUrl(url string) {
	//baseurl := "http://www.moe.gov.cn/jyb_xwfb/s271/"
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: "#list>li>a",
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
			TitleSelector: "#moe-detail-box>h1, #moe-detail-box>h2",
			TextSelector: ".TRS_Editor>p",
			DomainName: "http://www.moe.gov.cn/",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}

