package mohrss

import "testing"

func TestXpath(t *testing.T) {
	url := "http://www.mohrss.gov.cn/govsearch/xxgkjs_new2020.jsp?orderby=date&default=isall&page=45"
	GetPageUrlList(url)
	GetDetailPageUrl(url)
	infoUrl := "http://www.mohrss.gov.cn//xxgk2020/fdzdgknr/zcfg/bmgz/202011/t20201103_394781.html?keywords="
	infoUrl = "http://www.mohrss.gov.cn//xxgk2020/fdzdgknr/zcfg/gfxwj/rcrs/201407/t20140717_136235.html?keywords="
	infoUrl = "http://www.mohrss.gov.cn//xxgk2020/fdzdgknr/zcfg/gfxwj/rcrs/202105/t20210525_415135.html?keywords="
	GetHtmlInfo(infoUrl)

}