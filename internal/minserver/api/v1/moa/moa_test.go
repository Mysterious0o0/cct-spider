package moa

import "testing"

func TestXpath(t *testing.T) {

	url := "http://www.moa.gov.cn/gk/zcfg/qnhnzc/"
	GetPageUrlList(url)
	GetDetailPageUrl(url)
	infoUrl := "http://www.moa.gov.cn/govpublic/XZQYJ/202006/t20200617_6346579.htm"
	//infoUrl = "http://www.moa.gov.cn/govpublic/CWS/202012/t20201203_6357500.htm"
	//infoUrl = "http://www.moa.gov.cn/govpublic/CYZCFGS/202104/t20210426_6366592.htm"
	GetHtmlInfo(infoUrl)

}
