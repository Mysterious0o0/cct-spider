package miit

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/parse"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/urlprocess"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"hash"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

//var _cookie = _getCookie()

//func init()  {
//	go func() {
//		for range time.Tick(time.Minute * 5) {
//			_cookie = _getCookie()
//		}
//	}()
//}

func _getCookie(req request.Request, bCK string) (cookie string) {
	vm := otto.New()
	v, err := vm.Run(bCK)
	if err != nil {
		fmt.Println("otto run js error:", err)
		return
	}
	cookiePro := strings.Split(strings.Split(v.String(), "=")[1], ";")[0]
	ck := req.GetCookie("__jsluid_s")
	req.Cookies.StrCookie = fmt.Sprintf("%s; __jsl_clearance_s=%s", ck, cookiePro)

	b, _ := req.Visit()
	reg := regexp.MustCompile(`;go\((.*?)\)`)
	data := reg.FindStringSubmatch(string(b))
	c := _getjsluid(data[1])
	if c == "" {
		fmt.Println("getCookie error")
		return
	}
	cookie = fmt.Sprintf("%s; __jsl_clearance_s=%s", ck, c)
	return
}

func _getjsluid(ck string) string {
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

func GetPageUrlList(url string, urlChan chan<- *store.UrlChan, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		urlChan <- &store.UrlChan{
			Url:     fmt.Sprintf("%s%v", url, i),
			GetUrlF: GetDetailPageUrl,
		}

	}
}

func GetDetailPageUrl(url string, urlChan chan<- *store.UrlChan, infoChan chan<- *store.InfoChan) {
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	//req.Cookies.StrCookie = _cookie
	b, err := req.Visit()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	reg := regexp.MustCompile(`cookie=(\(.*?\));location`)
	cookie := reg.FindStringSubmatch(string(b))
	if len(cookie) > 0{
		req.Cookies.StrCookie = _getCookie(req, cookie[1])
		b, err = req.Visit()
	}
	var j store.JsonMiit
	err = json.Unmarshal(b, &j)
	if err != nil {
		logger.Error(err.Error(), logger.Field("url", url))
		return
	}
	for _, v := range j.DataMiit.DataResults {
		infoChan <- &store.InfoChan{
			Url:      urlprocess.UrlJoint(store.BaseUrl, v.GroupData[0].Url),
			GetInfoF: GetHtmlInfo,
		}
	}
}

func GetHtmlInfo(url string, errChan chan <- *store.InfoChan, message chan <- *store.Message){
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			Source: "工业和信息化部",
			DateSelector: "#con_time, .xxgk-span4",
			TitleSelector: "#con_title",
			TextSelector:  "#con_con>p",
			DomainName:    "https://www.miit.gov.cn/",
		},
	}
	//fmt.Println(_cookie)
	//pr.Request.Cookies.StrCookie = _cookie
	message <- pr.GetHtmlInfo()

	//infoMap := pr.GetHtmlInfo()
	//if len(infoMap) == 0 {
	//	errChan <- &store.InfoChan{
	//		Url:      url,
	//		GetInfoF: GetHtmlInfo,
	//	}
	//}else {
	//	info <- infoMap
	//}
}

func getPDFInfo(url string) (info []string) {
	req := request.Request{
		Url:    url,
		Method: http.MethodGet,
	}
	html, err := req.Visit()
	if err != nil {
		return
	}
	fmt.Println(string(html))
	return
}
