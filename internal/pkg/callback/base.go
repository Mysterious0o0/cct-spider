package callback

import (
	"sync"
)

type UrlChan struct {
	Url     string
	GetUrlF func(url string, urlChan chan<- *UrlChan, infoChan chan<- *InfoChan)
}

type InfoChan struct {
	Url      string
	GetInfoF func(url string, infoChan chan<- *InfoChan, message chan<- *Message)
}

type getUrlFunc func(url string, urlChan chan<- *UrlChan, infoChan chan<- *InfoChan)
type getInfoFunc func(url string, infoChan chan<- *InfoChan, message chan<- *Message)

func (f getUrlFunc) callUrl(url string, urlChan chan *UrlChan, infoChan chan *InfoChan) {
	f(url, urlChan, infoChan)
}

func (f getInfoFunc) callInfo(url string, infoChan chan *InfoChan, message chan *Message) {
	f(url, infoChan, message)
}

type callUrl interface {
	callUrl(string, chan *UrlChan, chan *InfoChan)
}

type callInfo interface {
	callInfo(string, chan *InfoChan, chan *Message)
}

func callback1(url string, c callUrl, urlChan chan *UrlChan, infoChan chan *InfoChan) {
	c.callUrl(url, urlChan, infoChan)
}

func callback2(url string, c callInfo, infoChan chan *InfoChan, message chan *Message) {
	c.callInfo(url, infoChan, message)
}

func (u UrlChan) GetUrlFunc(urlChan chan *UrlChan, infoChan chan *InfoChan, wg *sync.WaitGroup) {
	defer wg.Done()
	callback1(u.Url, getUrlFunc(u.GetUrlF), urlChan, infoChan)

}

func (u InfoChan) GetInfoFunc(infoChan chan *InfoChan, message chan *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	callback2(u.Url, getInfoFunc(u.GetInfoF), infoChan, message)
}

type Message struct {
	Url     string
	Title   string
	Content string
	Source  string
	Date    string
}

type SqlValues struct {
	NEWS_GUID        string // 新闻ID
	NEWS_TITLE       string // 新闻标题
	NEWS_TITLE_EN    string // 新闻标题英文
	NEWS_TS          string // 新闻时间戳
	NEWS_URL         string // 新闻URL
	NEWS_SOURCE      string // 新闻来源
	NEWS_SOURCE_CODE string // 新闻来源编码
	NEWS_SUMMARY     string // 新闻摘要
	POLICY_TYPE      string // 政策类型
	POLICY_TYPE_NAME string // 政策类型名称
	REGION_CODE      string // 地区代码
	REGION_NAME      string // 地区名称
	IS_CONTROL       string // 是否涉及出资企业
	IS_INVEST        string // 是否涉及投资企业
	IS_DEPOSIT       string // 是否涉及托管企业
	IS_FUND          string // 是否涉及基金企业
	IS_STOCK         string // 是否涉及股权企业
	IS_FINANCE       string // 是否涉及金融服务
	IS_INDUSTRY      string // 是否涉及实业控股
	IS_CAPITAL       string // 是否涉及资产经营
	NEWS_GYS_CODE    string // 新闻数据源代码
	NEWS_GYS_NAME    string // 新闻数据源名称
	NEWS_ID          int    // 新闻原始ID
}
