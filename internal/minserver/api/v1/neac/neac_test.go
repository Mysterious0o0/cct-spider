package neac

import "testing"

func TestXpath(t *testing.T) {

	url := "https://www.neac.gov.cn/seac/xxgk/zcjd/index.shtml"
	GetPageUrlList(url)
	GetDetailPageUrl(url)
	infoUrl := "https://www.neac.gov.cn/seac/xxgk/200703/1073907.shtml"
	infoUrl = "https://www.neac.gov.cn/seac/xxgk/201704/1073893.shtml"
	infoUrl = "https://www.neac.gov.cn/seac/xxgk/201108/1073902.shtml"
	GetHtmlInfo(infoUrl)
}

