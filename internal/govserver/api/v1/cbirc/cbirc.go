package cbirc

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/request"
	config2 "github.com/xiaogogonuo/cct-spider/pkg/config"
	"net/http"
)

func F() {
	c := config2.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/gov",
	}
	nc, err := c.NewConfig()
	if err != nil {
		return
	}
	req := request.Request{
		Url: nc.GetString("中国银行保险监督管理委员会"),
		Method: http.MethodGet,
	}
	b, err := req.Visit()
	fmt.Println(b, err)
}