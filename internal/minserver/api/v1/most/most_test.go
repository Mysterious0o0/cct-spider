package most

import "testing"

func TestXpath(t *testing.T) {
	//url := "http://www.most.gov.cn/kjzc/gjkjzc/"
	//GetFirstUrl(url)
	//pageUrl := "http://www.most.gov.cn/kjzc/gjkjzc/kyjggg/"
	//GetPageUrlList(pageUrl)
	//detailUrl := "http://www.most.gov.cn/kjzc/gjkjzc/kyjggg/"
	//GetDetailPageUrl(detailUrl)
	infoUrl := "http://www.most.gov.cn/xxgk/xinxifenlei/fdzdgknr/qtwj/qtwj2019/201901/t20190121_144852.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "http://www.most.gov.cn/tztg/200703/t20070309_41989.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "http://www.most.gov.cn/kjzc/gjkjzc/kyjggg/201703/t20170328_132207.html"
	GetHtmlInfo(infoUrl)
	infoUrl = "http://www.most.gov.cn/xxgk/xinxifenlei/fdzdgknr/fgzc/gfxwj/gfxwj2010before/201712/t20171222_137024.html"
	GetHtmlInfo(infoUrl)

}
