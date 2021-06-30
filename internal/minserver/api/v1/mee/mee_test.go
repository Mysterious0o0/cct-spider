package mee

import "testing"

func TestXpath(t *testing.T) {

	url := "http://www.mee.gov.cn/zcwj/"
	GetFirstUrl(url)
	firstUrl := "http://www.mee.gov.cn/zcwj/haqjwj/"
	firstUrl = "http://www.mee.gov.cn/zcwj/gwywj/"
	firstUrl = "http://www.mee.gov.cn/zcwj/bwj/"
	firstUrl = "http://www.mee.gov.cn/zcwj/bgtwj/"
	firstUrl = "http://www.mee.gov.cn/zcwj/xzspwj/"
	firstUrl = "http://www.mee.gov.cn/zcwj/haqjwj/"
	firstUrl = "http://www.mee.gov.cn/zcwj/qt/"
	GetSecondUrl(firstUrl)
	secondUrl := "http://www.mee.gov.cn/zcwj/bwj/wj/"
	secondUrl = "http://www.mee.gov.cn/zcwj/gwywj/"
	secondUrl = "http://www.mee.gov.cn/zcwj/xzspwj/"
	secondUrl = "http://www.mee.gov.cn/zcwj/bgtwj/han/"
	GetPageUrlList(secondUrl)
	pageUrl := "http://www.mee.gov.cn/zcwj/gwywj/index_23.shtml"
	pageUrl = "http://www.mee.gov.cn/zcwj/xzspwj/index_20.shtml"
	pageUrl = "http://www.mee.gov.cn/zcwj/bgtwj/han/index.shtml"
	GetDetailPageUrl(pageUrl)
	infoUrl := "http://www.mee.gov.cn/xxgk2018/xxgk/xxgk11/201806/t20180615_630017.html"
	infoUrl = "http://www.mee.gov.cn/zcwj/gwywj/201811/t20181129_676420.shtml"
	infoUrl = "http://www.mee.gov.cn/xxgk2018/xxgk/xxgk06/202105/t20210518_833282.html"
	GetHtmlInfo(infoUrl)
}