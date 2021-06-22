package mct

import "testing"

func TestXpath(t *testing.T) {

	//url := "http://zwgk.mct.gov.cn/zfxxgkml/447/458/463/index_3081.html"
	//GetPageUrlList(url)
	//GetDetailPageUrl(url)
	infoUrl := "http://zwgk.mct.gov.cn/zfxxgkml/zcfg/zcjd/202012/t20201205_915383.html"
	//infoUrl = "http://zwgk.mct.gov.cn/zfxxgkml/zcfg/zcjd/202105/t20210517_924523.html"
	//infoUrl = "http://zwgk.mct.gov.cn/zfxxgkml/zcfg/zcjd/202012/t20201205_915448.html"
	GetHtmlInfo(infoUrl)

}