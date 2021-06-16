package ndrc

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/respont"
	"net/http"
	"strings"
)

func GetFirstUrl(url string) {
	baseUrl := "https://www.ndrc.gov.cn/xxgk/"
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     baseUrl,
			UrlSelector: "ul[class='tab-menu moremenu'] a",
			Suffix:      "index.html",
		},
	}
	pr.GetPageUrl("href")
}

func GetPageUrlList(url string) {
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			PageNumSelector: ".page script",
		},
	}
	num := pr.GetPageNum()
	fmt.Println(num)
	if num == 0 {
		num = 20
	}
	fmt.Println(url) // first page url
	for i := 1; i < num; i++ {
		otherUrl := fmt.Sprintf("%s_%v.html", url[:len(url)-len(".html")], i)
		fmt.Println(otherUrl) // other page url
	}
}

func GetDetailPageUrl(url string) {
	s := strings.Split(url, "/")
	baseUrl := strings.Join(s[:len(s)-1], "/") + "/"
	pr := respont.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     baseUrl,
			UrlSelector: ".u-list a",
		},
	}
	pr.GetPageUrl("href")
}

func GetHtmlInfo(url string) (infoMap map[string]string) {
	infoMap = make(map[string]string)
	var pr respont.PR
	s := strings.Split(url, "/")
	//fmt.Println(s[5])
	r := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	switch s[5] {
	case "fzggwl":
		pr = respont.PR{
			Request: r,
			Parse: parse.Parse{
				TitleSelector: ".article_l>div div>span>strong, .article_l>div span>div>strong",
				TextSelector: ".TRS_Editor>p, .article_l>div>span",
			},
		}
	case "ghwb", "gg":
		pr = respont.PR{
			Request: r,
			Parse: parse.Parse{
				TitleSelector: ".TRS_Editor>div>strong>font>span, .TRS_Editor>div>span>strong, .TRS_Editor>span>div>strong, .TRS_Editor>div>font>span>strong, .TRS_Editor>p>strong>font",
				TextSelector: ".TRS_Editor>span, .TRS_Editor>p",
			},
		}
	case "qt", "tz", "ghxwj":
		pr = respont.PR{
			Request: r,
			Parse: parse.Parse{
				TitleSelector: ".article_l>div div>span>strong, .article_l>div span>div>strong",
				TextSelector: ".article_l span",
			},
		}
	}
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}
