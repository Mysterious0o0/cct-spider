package mof

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
)

func GetDetailPageUrl(url string) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: "#sub a",
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
			TitleSelector: ".testTxtH32",
			TextSelector: ".sqxzbList2 p",
			DomainName: "http://www.mof.gov.cn/",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}