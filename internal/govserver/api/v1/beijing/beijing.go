package beijing

import (
"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
"net/http"
	"sync"
)


func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {

	defer wg.Done()
	urlChan <- &callback.UrlChan{
		Url:     url,
		GetUrlF: GetDetailPageUrl,
	}
	//fmt.Println(url)
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: "body script",
		},
	}
	num := pr.GetPageNum("size:[0-9]+,")
	if num == 0 {
		num = 300
	}
	//fmt.Println(num)
	for i := 1; i < num; i++ {
		urlChan <- &callback.UrlChan{
			Url:     fmt.Sprintf("%sindex_%v.html", url, i),
			GetUrlF: GetDetailPageUrl,
		}
		//fmt.Printf("%sindex_%v.html\n", url, i)
	}
}

func GetDetailPageUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {

	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			UrlSelector: ".list a",
			BaseUrl: url,
		},
	}
	urlList := pr.GetPageUrl("href")
	for _, link := range urlList {
		//fmt.Println(link)
		infoChan <- &callback.InfoChan{
			Url:      link,
			GetInfoF: GetHtmlInfo,
		}
	}
}

func GetHtmlInfo(url string, errChan chan<- *callback.InfoChan, message chan<- *callback.Message) {

	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			Source:        "北京市人民政府",
			SourceCode:    "WEB_01132",
			DateSelector:  "ol[class='doc-info clearfix']>li:nth-child(8)>span",
			TitleSelector: "h1",
			TextSelector:  "#mainText",
			DomainName:    "www.beijing.gov.cn",

		},
	}
	message <- pr.GetHtmlInfo()

	//infoMap := pr.GetHtmlInfo()
	//if len(infoMap) == 0 {
	//	errChan <- &store.InfoChan{
	//		Url:      url,
	//		GetInfoF: GetHtmlInfo,
	//	}
	//}else {
	//	info <- infoMap
	//}
}


