package mps

import "testing"

func TestXpath(t *testing.T) {

	url := "https://www.mps.gov.cn/n6557558/index.html"
	GetPageUrlList(url)
	url = "https://www.mps.gov.cn/n6557558/index_7574611_14.html"
	GetDetailPageUrl(url)
	infoUrl := "https://www.mps.gov.cn/n6557558/c7789058/content.html"
	infoUrl = "https://www.mps.gov.cn/n6557558/c3618764/content.html"
	infoUrl = "https://www.mps.gov.cn/n6557558/c6836853/content.html"
	GetHtmlInfo(infoUrl)
}
