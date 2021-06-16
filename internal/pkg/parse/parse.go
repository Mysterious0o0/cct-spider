package parse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Parse struct {
	Html            string
	BaseUrl         string
	UrlSelector     string
	TitleSelector   string
	TextSelector    string
	PageNumSelector string
	Suffix          string
}

func (p Parse) GetTextByParseHtml() (title string, info []string) {
	var titleL []string
	var key map[string]byte
	key = make(map[string]byte)
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		log.Fatalln(err)
		return
	}
	dom.Find(p.TitleSelector).Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != "" {
			titleL = append(titleL, selection.Text())
		}
	})
	title = strings.Join(titleL, "")

	dom.Find(p.TextSelector).Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		if _, ok := key[s]; s != "" && !ok{
			info = append(info, s)
			key[s] = 0
		}
	})
	return
}

func (p Parse) GetOneUrlByParseHtml(attrName string) (src string, b bool) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		log.Fatalln(err)
		return
	}
	dom.Find(p.UrlSelector).Each(func(i int, selection *goquery.Selection) {
		src, b = selection.Attr(attrName)
	})
	return
}

func (p Parse) GetAllUrlByParseHtml(attrName string) (hrefList []string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		log.Fatalln(err)
		return
	}
	dom.Find(p.UrlSelector).Each(func(i int, selection *goquery.Selection) {
		href, b := selection.Attr(attrName)
		if b && href != "" {
			if strings.Contains(href, "http") || strings.Contains(href, "https"){
				hrefList = append(hrefList, href + p.Suffix)
				fmt.Println(href + p.Suffix)
			}else {
				hrefList = append(hrefList, p.BaseUrl + href +p.Suffix)
				fmt.Println(p.BaseUrl + href +p.Suffix)
			}
		}else {
			log.Fatalln(b, href)
		}
	})
	return
}


func (p Parse) GetPageNum(r string) (num int) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		log.Fatalln(err)
		return
	}
	dom.Find(p.PageNumSelector).Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != "" {
			reg := regexp.MustCompile(r)
			numStr := reg.FindString(selection.Text())
			if numStr == "" {
				return
			}
			num, err = strconv.Atoi(numStr)
			if err != nil {
				log.Fatalln(err)
			}
		}
	})
	return
}

func (p Parse) GetCountAndSize(countR string, sizeR string) (count int, size int) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(p.Html))
	if err != nil {
		log.Fatalln(err)
		return
	}
	dom.Find(p.PageNumSelector).Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		if s != "" {
			countReg := regexp.MustCompile(countR)
			sizeReg := regexp.MustCompile(sizeR)
			numReg := regexp.MustCompile("\\d+")
			countStr := countReg.FindString(s)
			countStr = numReg.FindString(countStr)
			sizeStr := sizeReg.FindString(s)
			sizeStr = numReg.FindString(sizeStr)
			if countStr != ""{
				count, err = strconv.Atoi(countStr)
				if err != nil {
					log.Fatalln(err)
				}
			}else {
				fmt.Println("count is nil")
			}

			if sizeStr != ""{
				size, err = strconv.Atoi(sizeStr)
				if err != nil {
					log.Fatalln(err)
				}
			}else {
				fmt.Println("size is nil")
			}
		}
	})
	return
}