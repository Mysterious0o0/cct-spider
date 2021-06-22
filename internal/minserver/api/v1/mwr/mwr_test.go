package mwr

import (
	"testing"
)

func TestXpath(t *testing.T) {

	//url := "http://www.mwr.gov.cn/zw/zcjd/"
	//GetPageUrlList(url)
	//GetDetailPageUrl(url)
	infoUrl := "http://www.mwr.gov.cn/hd/zxft/zxzb/fbh20200821/"
	infoUrl = "http://www.mwr.gov.cn/zw/zcjd/202009/t20200930_1449711.html"
	infoUrl = "http://www.mwr.gov.cn/zw/zcjd/202004/t20200408_1399443.html"
	GetHtmlInfo(infoUrl)

}
