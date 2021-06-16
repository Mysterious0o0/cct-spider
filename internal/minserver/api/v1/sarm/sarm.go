package miit

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/respont"
	"net/http"
)

func GetFirstUrl(url string) {
	fmt.Println(url) // frist url
	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("%s_%v.html", url[:len(url)-len(".html")], i)
		fmt.Println(url) // add other url
	}
}

func GetDetailPageUrl(url string) {
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			UrlSelector: ".Three_zhnlist_02 a",
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
	//	UrlSelector: ".Three_zhnlist_02 a",
	//}
	//hrefL := p.GetAllUrlByParseHtml("href")
	//for _, href := range hrefL{
	//	fmt.Println(href)
	//}
}

func GetHtmlInfo(url string) (infoMap map[string]string){
	infoMap = make(map[string]string)
	pr := respont.PR{
		Request: request.Request{
			Url : url,
			Method: http.MethodGet,
		},
		Parse:   parse.Parse{
			TitleSelector: ".xilanboxbg td[colspan='2'] li[class='Three_xilan01_02 Three_xilan01_0201']",
			TextSelector: ".Three_xilan_07 p",
		},
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}
