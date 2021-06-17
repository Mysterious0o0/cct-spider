package fmprc

import (
	"testing"
)

func TestXpath(t *testing.T) {
	url := "https://www.fmprc.gov.cn/web/wjb_673085/zfxxgk_674865/gknrlb/zcfg/"
	GetDetailPageUrl(url)
	infoUrl := "https://www.fmprc.gov.cn/web/wjb_673085/zfxxgk_674865/gknrlb/zcfg/./gfxwj/t1723270.shtml"
	//infoUrl = "https://www.fmprc.gov.cn/web/wjb_673085/zfxxgk_674865/gknrlb/zcfg/./flfg/t70823.shtml"
	//infoUrl = "http://www.caac.gov.cn/XXGK/XXGK/TZTG/202007/t20200721_203688.html"
	GetHtmlInfo(infoUrl)

}


