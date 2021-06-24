package cbirc

import (
	"encoding/json"
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"net/http"
	"strings"
	"sync"
)

func GetPageUrlList(url string, infoChan chan<- *store.InfoChan, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		req := request.Request{
			Url:    fmt.Sprintf(url, i),
			Method: http.MethodGet,
		}
		b, err := req.Visit()
		if err != nil {
			break
		}
		var j store.JsonCbirc
		err = json.Unmarshal(b, &j)
		if err != nil {
			fmt.Println(err)
		}
		if len(j.Data.Rows) == 0 {
			break
		}
		for _, v := range j.Data.Rows {
			//fmt.Printf(store.DetailUrl, v.DocId)
			infoChan <- &store.InfoChan{
				Url:      fmt.Sprintf(store.DetailUrl, v.DocId),
				GetInfoF: GetHtmlInfo,
			}
		}
	}
}


func GetHtmlInfo(url string, errChan chan <- *store.InfoChan, info chan <-map[string]string) {
	infoMap := make(map[string]string)
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	if err != nil {
		return
	}
	var j store.JsonDetailsCbirc
	err = json.Unmarshal(b, &j)
	if err != nil {
		fmt.Println(err)
	}

	p := parse.Parse{
		Html: j.DocClob,
		TextSelector:  "p",
	}
	_, data := p.GetTextByParseHtml()
	infoMap[j.DocTitle] = strings.Join(data, "")
	info <- infoMap
	//if len(infoMap) == 0 {
	//	errChan <- &store.InfoChan{
	//		Url:      url,
	//		GetInfoF: GetHtmlInfo,
	//	}
	//}else {
	//	info <- infoMap
	//}
}