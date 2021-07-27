package zhejiang

import (
	"encoding/json"
	"fmt"
	"github.com/489397771/cct-spider/internal/govserver/store"
	"github.com/489397771/cct-spider/internal/pkg/callback"
	"github.com/489397771/cct-spider/internal/pkg/parse"
	"github.com/489397771/cct-spider/internal/pkg/request"
	"github.com/489397771/cct-spider/internal/pkg/response"
	"github.com/489397771/cct-spider/pkg/logger"
	"net/http"
	"strings"
)


//func GetPageUrlList(url string, urlChan chan<- *callback.UrlChan, wg *sync.WaitGroup) {

func GetPageUrlList(url string) {

	//defer wg.Done()
	req := request.Request{
		Url: url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	if err != nil{
		return
	}
	var j store.ZJJson
	err = json.Unmarshal(b, &j)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}
	req.Url = strings.Replace(url, `page_size=1`, fmt.Sprintf(`page_size=%v`, j.Total), -1)
	b, err = req.Visit()
	if err != nil{
		return
	}
	err = json.Unmarshal(b, &j)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}
	for _, v := range j.Data {
		fmt.Println(v.Url)
		//infoChan <- &callback.InfoChan{
		//	Url:      v.Url,
		//	GetInfoF: GetHtmlInfo,
		//}
	}
}


func GetHtmlInfo(url string, errChan chan<- *callback.InfoChan, message chan<- *callback.Message) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			Source:        "广州市人民政府",
			SourceCode:    "WEB_01187",
			DateSelector:  "div[class='classify']>table>tbody>:nth-child(4)>td[class='td-value']>span",
			TitleSelector: ".content>h1",
			TextSelector:  ".article-content>center",
			DomainName:    "www.gd.gov.cn",

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