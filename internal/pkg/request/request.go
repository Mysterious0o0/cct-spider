package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Url     string
	Method  string
	Timeout time.Duration
	Body    io.Reader
	Param   map[string]string
	Header  map[string]string
	Cookies struct {
		StrCookie  string
		HttpCookie []*http.Cookie
	}
}

func (r *Request) Visit() (b []byte, err error) {
	resp, err := r.request()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return
}

func (r *Request) GetCookie(name string) (ck string) {
	var cks []string
	for _, cookie := range r.Cookies.HttpCookie {
		if cookie.Name == name {
			cks = append(cks, cookie.String())
		}
	}
	ck = strings.Join(cks, "")
	return
}

func (r *Request) request() (resp *http.Response, err error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36")
	q := req.URL.Query()
	for key, val := range r.Param {
		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	if r.Cookies.StrCookie != ""{
		req.Header.Set("Cookie", r.Cookies.StrCookie)
	}
	client := &http.Client{Timeout: r.Timeout}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != 521 {
		fmt.Println(resp.StatusCode, r.Url)
		return
	}
	r.Cookies.HttpCookie = resp.Cookies()
	return
}
