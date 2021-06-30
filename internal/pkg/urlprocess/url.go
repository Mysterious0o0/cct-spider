package urlprocess

import (
	"fmt"
	"net/url"
)

func UrlJoint(baseUrl string, tailUrl string) string{
	u, err := url.Parse(tailUrl)
	if err != nil {
		fmt.Println(err)
	}
	base, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(base.ResolveReference(u))
	return base.ResolveReference(u).String()
}
