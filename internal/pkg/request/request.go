package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
	Url     string
	Method  string
	Timeout time.Duration
	Body    io.Reader
	Param   map[string]string
	Header  map[string]string
	Cookie  string
}

func (r *Request) Visit() (b []byte, err error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent" ,"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36")
	q := req.URL.Query()
	for key, val := range r.Param {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	if r.Cookie != ""{
		req.Header.Set("Cookie", r.Cookie)
	}
	client := &http.Client{Timeout: r.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != 521 {
		fmt.Println(resp.StatusCode, r.Url)
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return
}
