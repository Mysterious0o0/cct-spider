package mohrss

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
)

func GetPageUrlList(url string) {
	num := 60
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: "body script",
		},
	}
	countR := "var m_nRecordCount = \\d+"
	sizeR := "var m_nPageSize = \\d+"
	count, size := pr.GetCountAndSize(countR, sizeR)
	fmt.Println(count, size)
	if count != 0 && size != 0 {
		num = count / size + 1
	}
	for i := 1; i <= num; i++ {
		otherUrl := fmt.Sprintf("%s%v", url[:len(url)-1], i)
		fmt.Println(otherUrl) // all page url
	}
}

func GetDetailPageUrl(url string) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			UrlSelector: "tbody tr a",
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
			DomainName: "http://www.mohrss.gov.cn/",
			TitleSelector: ".artT",
			TextSelector: ".Custom_UnionStyle p, .TRS_PreAppend p, div[class='art_p leftW'] p, .art_p p",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}