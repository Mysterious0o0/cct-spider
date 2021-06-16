package respont

import (
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"strings"
)

type PR struct {
	Request request.Request
	Parse   parse.Parse
}
func (pr PR)GetPageUrl(attrName string) {
	html, err := pr.Request.Visit()
	if err != nil {
		return
	}
	pr.Parse.Html = string(html)
	pr.Parse.GetAllUrlByParseHtml(attrName)
}

func (pr PR)GetHtmlInfo() (infoMap map[string]string){
	var info []string
	infoMap = make(map[string]string)
	html, _ := pr.Request.Visit()
	pr.Parse.Html = string(html)
	title, data := pr.Parse.GetTextByParseHtml()
	info = append(info, data...)
	if _, ok := infoMap[title]; !ok{
		infoMap[title] = strings.Join(info, "")
	}
	return
}

func (pr PR)GetPageNum(r string) (num int){
	html, _ := pr.Request.Visit()
	pr.Parse.Html = string(html)
	num = pr.Parse.GetPageNum(r)
	return
}


func (pr PR)GetCountAndSize(countR string, sizeR string) (count int, size int){
	html, _ := pr.Request.Visit()
	pr.Parse.Html = string(html)
	count, size = pr.Parse.GetCountAndSize(countR, sizeR)
	return
}

