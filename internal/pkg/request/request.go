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
	body    io.Reader
	Param   map[string]string
	Header  map[string]string
	Cookie  map[string]string
}

func (r *Request) Visit() (b []byte, err error) {
	req, err := http.NewRequest(r.Method, r.Url, r.body)
	if err != nil {
		return
	}
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
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return
}
