package mps

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"hash"
	"net/http"
	"regexp"
	"strings"
)

var _cookie = GetCookie()

func GetCookie() (cookie string) {
	url := "https://www.mps.gov.cn/n6557558/index.html"
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	b, _ := req.Visit()
	reg := regexp.MustCompile(`cookie=(\(.*?\));location`)
	jslClearances := reg.FindStringSubmatch(string(b))
	vm := otto.New()
	v, err := vm.Run(jslClearances[1])
	if err != nil {
		fmt.Println("otto run js error:", err)
		return
	}
	cookiePro := strings.Split(strings.Split(v.String(), "=")[1], ";")[0]
	ck := req.GetCookie("__jsluid_s")
	req.Cookies.StrCookie = fmt.Sprintf("%s; __jsl_clearance_s=%s", ck, cookiePro)

	b, _ = req.Visit()
	reg = regexp.MustCompile(`;go\((.*?)\)`)
	data := reg.FindStringSubmatch(string(b))
	c := _getCookie(data[1])
	if c == "" {
		fmt.Println("getCookie error")
		return
	}
	cookie = fmt.Sprintf("%s; __jsl_clearance_s=%s", ck, c)
	return
}

func _getCookie(ck string) string {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(ck), &data)
	if err != nil {
		fmt.Println("[]byte -> map err,", err)
		return ""
	}
	chars := fmt.Sprintf("%s", data["chars"].(string))
	charsLen := len(chars)
	for i := 0; i < charsLen; i++ {
		for j := 0; j < charsLen; j++ {
			clearance := fmt.Sprintf("%s%c%c%s",
				data["bts"].([]interface{})[0].(string), chars[i], chars[j], data["bts"].([]interface{})[1].(string))
			var encrypt hash.Hash
			switch data["ha"].(string) {
			case "md5":
				encrypt = md5.New()
			case "sha1":
				encrypt = sha1.New()
			default:
				encrypt = sha256.New()
			}
			encrypt.Write([]byte(clearance))
			s := hex.EncodeToString(encrypt.Sum(nil))
			if s == data["ct"].(string) {
				return clearance
			}
		}
	}
	return ""
}

func GetFirstUrl(url string) {
	fmt.Println(url) // frist url
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: "div[style='display:none'] a",
		},
	}
	pr.Request.Cookies.StrCookie = _cookie
	pr.GetPageUrl("href")
}

func GetDetailPageUrl(url string) {
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			BaseUrl:     url,
			UrlSelector: ".list a",
		},
	}
	pr.Request.Cookies.StrCookie = _cookie
	pr.GetPageUrl("href")
}

func GetHtmlInfo(url string) (infoMap map[string]string) {
	infoMap = make(map[string]string)
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			DomainName:    "https://www.mps.gov.cn/",
			TitleSelector: "div[class='bTitle w915']>p",
			TextSelector:  "div[class='wordContent w915'] p",
		},
	}
	pr.Request.Cookies.StrCookie = _cookie
	infoMap = pr.GetHtmlInfo()
	fmt.Println(infoMap)
	return
}
