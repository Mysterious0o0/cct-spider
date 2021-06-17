package respont

import (
	"fmt"
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
	if !strings.Contains(pr.Request.Url, pr.Parse.DomainName) {
		fmt.Printf("域名：%s 网址：%s 域名不存在\n", pr.Request.Url, pr.Parse.DomainName)
	}
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

