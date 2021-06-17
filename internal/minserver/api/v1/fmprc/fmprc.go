package fmprc

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/respont"
	"net/http"
)

func GetDetailPageUrl(url string, baseUrl string) {
	baseUrl = "https://www.fmprc.gov.cn/web/wjb_673085/zfxxgk_674865/gknrlb/zcfg/"
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     baseUrl,
			UrlSelector: "div[class='jgbox_r fr'] a[target='_blank']",
		},
	}
	pr.GetPageUrl("href")

	//req := request.Request{
	//	Url:    url,
	//	Method: http.MethodGet,
	//}
	//html, err := req.Visit()
	//if err != nil {
	//	return
	//}
	//p := parse.Parse{
	//	Html:        string(html),
	//	UrlSelector: "div[class='jgbox_r fr'] a[target='_blank']",
	//}
	//hrefL := p.GetAllUrlByParseHtml("href")
	//for _, href := range hrefL{
	//	if strings.Contains(href, "http") || strings.Contains(href, "https"){
	//		fmt.Println(href)
	//	}else {
	//		fmt.Println(baseUrl + href)
	//	}
	//}
}

func GetHtmlInfo(url string) (infoMap map[string]string) {
	infoMap = make(map[string]string)
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			TitleSelector: "#News_Body_Title",
			TextSelector:  "#News_Body_Txt_A>p",
			DomainName:    "https://www.fmprc.gov.cn/",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}
