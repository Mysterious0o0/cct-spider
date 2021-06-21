package mnr

import "testing"

func TestXpath(t *testing.T) {
	url := "http://f.mnr.gov.cn/index_3553.html"
	GetFirstUrl(url)
	pageUrl := "http://f.mnr.gov.cn/index_3553_80.html"
	GetDetailPageUrl(pageUrl)
	infoUrl := "http://f.mnr.gov.cn/202104/t20210406_2619380.html"
	infoUrl = "http://f.mnr.gov.cn/202103/t20210326_2618380.html"
	infoUrl = "http://f.mnr.gov.cn/201702/t20170206_1436110.html"
	GetHtmlInfo(infoUrl)
}
