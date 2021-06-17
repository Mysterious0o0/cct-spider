package ndrc

import (
	"testing"
)

func TestXpath(t *testing.T) {
	url := "https://www.ndrc.gov.cn/xxgk/"
	GetFirstUrl(url)
	pageUrl := "https://www.ndrc.gov.cn/xxgk/./zcfb/gg/index.html"
	GetPageUrlList(pageUrl)
	detailUrl := "https://www.ndrc.gov.cn/xxgk/./zcfb/gg/index_1.html"
	GetDetailPageUrl(detailUrl)
	infoUrl := "https://www.ndrc.gov.cn/xxgk/zcfb/fzggwl/201409/t20140915_960787.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "https://www.ndrc.gov.cn/xxgk/zcfb/ghxwj/202101/t20210119_1265242.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "https://www.ndrc.gov.cn/xxgk/zcfb/gg/201904/t20190412_961217.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "https://www.ndrc.gov.cn/xxgk/zcfb/ghwb/201908/t20190821_962257.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "https://www.ndrc.gov.cn/xxgk/zcfb/tz/202105/t20210506_1279236.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "https://www.ndrc.gov.cn/xxgk/zcfb/qt/202105/t20210527_1281209.html"
	GetHtmlInfo(infoUrl)
}

//func TestGo(t *testing.T) {
//	detailUrl := "https://www.ndrc.gov.cn/xxgk/zcfb/qt/202105/t20210527_1281206.html"
//	var wg sync.WaitGroup
//	for i := 0; i < 100; i++{
//		wg.Add(1)
//		go GetHtmlInfo(detailUrl, &wg)
//	}
//	wg.Wait()
//}
