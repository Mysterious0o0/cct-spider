package request

import (
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
	Cookie  map[string]string
}

func (r *Request) Visit() (b []byte, err error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent" ,"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	q := req.URL.Query()
	for key, val := range r.Param {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()
	for k, v := range r.Header {
		req.Header.Add(k, v)
	}
	for k, v := range r.Cookie {
		req.AddCookie(&http.Cookie{
			Name: k,
			Value: v,
		})
	}
	client := &http.Client{Timeout: r.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return
}
