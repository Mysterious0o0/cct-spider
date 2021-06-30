package miit

import (
	"fmt"
	"testing"
	"time"
)

//func TestXpath(t *testing.T) {
//	url := "https://www.miit.gov.cn/search-front-server/api/search/info?websiteid=110000000000000&scope=basic&q=&pg=15&cateid=57&pos=title_text%2Cinfocontent%2Ctitlepy&selectFields=title,url,&group=distinct&level=6&sortFields=%5B%7B%22name%22%3A%22deploytime%22%2C%22type%22%3A%22desc%22%7D%5D&p="
//	GetPageUrlList(url)
//	pageUrl := "https://www.miit.gov.cn/search-front-server/api/search/info?websiteid=110000000000000&scope=basic&q=&pg=15&cateid=57&pos=title_text%2Cinfocontent%2Ctitlepy&selectFields=title,url,&group=distinct&level=6&sortFields=%5B%7B%22name%22%3A%22deploytime%22%2C%22type%22%3A%22desc%22%7D%5D&p=1"
//	GetDetailPageUrl(pageUrl)
//	infoUrl := "https://www.miit.gov.cn/zwgk/zcwj/wjfb/qt/art/2021/art_9b8e1e5711c54cec9ca0cda73ac36c40.html"
//	infoUrl = "https://www.miit.gov.cn/zwgk/zcwj/wjfb/tg/art/2021/art_b5fbf1361d314615a63f8dae9646c511.html"
//	GetHtmlInfo(infoUrl)
//
//}

func TestA(t *testing.T)  {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Minute * 1)
		fmt.Println(_cookie)
	}
}