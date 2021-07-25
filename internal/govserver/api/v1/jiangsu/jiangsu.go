package jiangsu

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"net/http"
)

//var baseUrl = `https://www.ah.gov.cn/site/label/8888?IsAjax=1&labelName=maYaService&pageSize=13&pageIndex=%v&domain=http%%3A%%2F%%2F192.168.60.62%%3A8090%%2Fxxgk%%2Fsitesearch%%3FpageSize%%3D13%%26page%%3D%v%%26pageIndex%%3D%v%%26scopeType%%3D1%%26matchType%%3D1%%26rangeType%%3D1%%26classname%%3D%%26documentType%%3D%%26dsId%%3Dwww.ah.gov.cn%%26q%%3D%%26t%%3D1627027993128&file=%%2Fahxxgk%%2Fxxgk%%2FpublicInfoList_new2020_ah_node6`

//func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {

func GetPageUrlList(url string) {
	//defer wg.Done()
	//urlChan <- &callback.UrlChan{
	//	Url:     url,
	//	GetUrlF: GetDetailPageUrl,
	//}
	fmt.Println(url)
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
		//urlChan <- &callback.UrlChan{
		//	Url:     fmt.Sprintf(baseUrl, i, i, i),
		//	GetUrlF: GetDetailPageUrl,
		//}
		fmt.Printf("%sindex_%v.html\n", url, i)
	}
}

//func GetDetailPageUrl(url string, urlChan chan<- *callback.UrlChan, infoChan chan<- *callback.InfoChan) {

func GetDetailPageUrl(url string) {

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
		fmt.Println(link)

		//infoChan <- &callback.InfoChan{
		//	Url:      link,
		//	GetInfoF: GetHtmlInfo,
		//}
	}
}

//func GetHtmlInfo(url string, errChan chan<- *callback.InfoChan, message chan<- *callback.Message) {

func GetHtmlInfo(url string) {
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
	//message <- pr.GetHtmlInfo()
	m := pr.GetHtmlInfo()
	fmt.Println(m.Summary)
	fmt.Println("----------------------")
	fmt.Println(m.Title)
	fmt.Println("----------------------")
	fmt.Println(m.Date)

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
