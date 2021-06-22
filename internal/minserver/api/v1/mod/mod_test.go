package mod

import "testing"

func TestXpath(t *testing.T) {
	url := "http://www.mod.gov.cn/regulatory/node_47883.htm"
	GetPageUrlList(url)
	GetDetailPageUrl(url)
	infoUrl := "http://www.mod.gov.cn/regulatory/2015-07/03/content_4643973.htm"
	infoUrl = "http://www.mod.gov.cn/regulatory/2017-10/10/content_4794244.htm"
	infoUrl = "http://www.mod.gov.cn/regulatory/2017-07/17/content_4785976.htm"
	GetHtmlInfo(infoUrl)

}
