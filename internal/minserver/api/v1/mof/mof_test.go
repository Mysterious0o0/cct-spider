package mof

import "testing"


func TestXpath(t *testing.T) {
	pageUrl := "http://www.mof.gov.cn/gp/xxgkml/index_8254.htm"
	GetDetailPageUrl(pageUrl)
	infoUrl := "http://www.mof.gov.cn/gp/xxgkml/kjs/202106/t20210618_3721801.htm"
	infoUrl = "http://www.mof.gov.cn/gp/xxgkml/gks/202009/t20200911_3587105.htm"
	infoUrl = "http://www.mof.gov.cn/gp/xxgkml/gks/202009/t20200924_3594315.htm"
	infoUrl = "http://www.mof.gov.cn/gp/xxgkml/gks/202008/t20200828_3577004.htm"
	GetHtmlInfo(infoUrl)
}