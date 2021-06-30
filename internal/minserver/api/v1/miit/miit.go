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
	"hash"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var _cookie = _getCookie()

func init()  {
	go func() {
		for range time.Tick(time.Minute * 1) {
			_cookie = _getCookie()
		}
	}()
}

func _getCookie() (cookie string) {
	url := "https://www.miit.gov.cn/search-front-server/api/search/info"
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
	//fmt.Println(_cookie)
	req.Cookies.StrCookie = _cookie
	b, err := req.Visit()
	//fmt.Println(string(b))
	if err != nil {
		return
	}
	var j store.JsonMiit
	err = json.Unmarshal(b, &j)
	if err != nil {
		fmt.Println(url, err)
		return
	}
	//fmt.Println(j.Data.Total) //总文件数
	for _, v := range j.DataMiit.DataResults {
		//fmt.Println(urlprocess.UrlJoint(store.BaseUrl, v.GroupData[0].Url))
		//fmt.Println(baseUrl + v.GroupData[0].Url)
		infoChan <- &store.InfoChan{
			Url:      urlprocess.UrlJoint(store.BaseUrl, v.GroupData[0].Url),
			GetInfoF: GetHtmlInfo,
		}
	}
}

func GetHtmlInfo(url string, errChan chan <- *store.InfoChan, info chan <-map[string]string){
	pr := response.PR{
		Request: request.Request{
			Url:    url,
			Method: http.MethodGet,
		},
		Parse: parse.Parse{
			//UrlSelector:   "iframe",
			TitleSelector: "#con_title",
			TextSelector:  "#con_con>p",
			DomainName:    "https://www.miit.gov.cn/",
		},
	}
	//fmt.Println(_cookie)
	pr.Request.Cookies.StrCookie = _cookie
	info <- pr.GetHtmlInfo()

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
