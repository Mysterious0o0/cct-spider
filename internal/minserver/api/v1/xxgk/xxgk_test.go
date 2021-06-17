package miit

import (
	"testing"
)

func TestXpath(t *testing.T) {

	url := "https://xxgk.mot.gov.cn/2020/jigou/list.html"
	GetFirstUrl(url)
	GetDetailPageUrl(url)
	infoUrl := "https://xxgk.mot.gov.cn/2020/jigou/./fgs/202106/t20210604_3605909.html"
	infoUrl = "https://xxgk.mot.gov.cn/2020/jigou/./ysfws/202106/t20210607_3607600.html"
	infoUrl = "https://xxgk.mot.gov.cn/2020/jigou/./aqyzljlglj/202106/t20210611_3609694.html"
	GetHtmlInfo(infoUrl)
}


