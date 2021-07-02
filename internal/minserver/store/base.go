package store

import (
	"sync"
)

type UrlChan struct {
	Url     string
	GetUrlF func(s string, urlChan chan<- *UrlChan, infoChan chan<- *InfoChan)
}

type InfoChan struct {
	Url      string
	GetInfoF func(s string, infoChan chan<- *InfoChan, info chan<- *Message)
}

type getUrlFunc func(url string, urlChan chan<- *UrlChan, infoChan chan<- *InfoChan)
type getInfoFunc func(url string, infoChan chan<- *InfoChan, info chan<- *Message)

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
