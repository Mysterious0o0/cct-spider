package moe

import "testing"

func TestXpath(t *testing.T) {
	url := "http://www.moe.gov.cn/jyb_xwfb/s271/"
	GetPageUrlList(url)
	GetDetailPageUrl(url)
	infoUrl := "http://www.moe.gov.cn/jyb_xwfb/s271/202106/t20210601_534685.html"
	infoUrl = "http://www.moe.gov.cn/jyb_xwfb/s271/./202104/t20210409_525364.html"
	infoUrl = "http://www.moe.gov.cn/jyb_xwfb/s271/./202101/t20210115_509949.html"
	GetHtmlInfo(infoUrl)

}