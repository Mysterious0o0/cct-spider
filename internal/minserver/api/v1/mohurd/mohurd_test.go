package mohurd

import "testing"

func TestXpath(t *testing.T) {
	url := "http://www.mohurd.gov.cn/index.html"
	GetFirstUrl(url)
	firstUrl := "http://www.mohurd.gov.cn/fgjs/index.html"
	firstUrl = "http://www.mohurd.gov.cn/zfggyfz/index.html"
	firstUrl = "http://www.mohurd.gov.cn/zfbz/index.html"
	firstUrl = "http://www.mohurd.gov.cn/cxgh/index.html"
	firstUrl = "http://www.mohurd.gov.cn/bzde/index.html"
	firstUrl = "http://www.mohurd.gov.cn/fdcy/index.html"
	firstUrl = "http://www.mohurd.gov.cn/jzsc/index.html"
	firstUrl = "http://www.mohurd.gov.cn/csjs/index.html"
	firstUrl = "http://www.mohurd.gov.cn/czjs/index.html"
	firstUrl = "http://www.mohurd.gov.cn/zlaq/index.html"
	firstUrl = "http://www.mohurd.gov.cn/jzjnykj/index.html"
	firstUrl = "http://www.mohurd.gov.cn/zfgjjjg/index.html"
	firstUrl = "http://www.mohurd.gov.cn/csgljd/index.html"
	firstUrl = "http://www.mohurd.gov.cn/xytj/index.html"
	firstUrl = "http://www.mohurd.gov.cn/gjjlyhz/index.html"
	firstUrl = "http://www.mohurd.gov.cn/cjda/index.html"
	firstUrl = "http://www.mohurd.gov.cn/jsrc/index.html"
	firstUrl = "http://www.mohurd.gov.cn/djgz/index.html"
	firstUrl = "http://www.mohurd.gov.cn/ltxgbgz/index.html"
	firstUrl = "http://www.mohurd.gov.cn/xwfbzxft/index.html"
	firstUrl = "http://www.mohurd.gov.cn/zcjd/index.html"
	firstUrl = "http://www.mohurd.gov.cn/gzly/index.html"
	firstUrl = "http://www.mohurd.gov.cn/xf/index.html"
	GetSecondUrl(firstUrl)
	secondUrl := "http://www.mohurd.gov.cn/fgjs/fl/index.html"
	secondUrl = "http://www.mohurd.gov.cn/fdcy/fdcyxydt/index.html"
	secondUrl = "http://www.mohurd.gov.cn/csjs/csjsdfxx/index.html"
	secondUrl = "http://www.mohurd.gov.cn/zlaq/zlaqdfxx/index.html"
	GetPageUrlList(secondUrl)
	pageUrl := "http://www.mohurd.gov.cn/fgjs/fl/index_3.html"
	pageUrl = "http://www.mohurd.gov.cn/zlaq/zlaqdfxx/index_38.html"
	pageUrl = "http://www.mohurd.gov.cn/csjs/csjsdfxx/index_31.html"
	GetDetailPageUrl(pageUrl)
	infoUrl := "http://www.mohurd.gov.cn/dfxx/202001/t20200113_243531.html"
	infoUrl = "http://www.mohurd.gov.cn/dfxx/201702/t20170222_230688.html"
	infoUrl = "http://www.mohurd.gov.cn/fgjs/fl/200611/t20061101_159473.html"
	GetHtmlInfo(infoUrl)
}
