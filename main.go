package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/cbirc"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/mee"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/miit"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/sarm"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	M "github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"sync"
)

func minConfig() *viper.Viper {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/min",
	}
	v, err := c.NewConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func ministries() {

	wg := &sync.WaitGroup{}
	urlChannel := make(chan *store.UrlChan)
	infoChannel := make(chan *store.InfoChan)
	errChannel := make(chan *store.InfoChan)
	infoMap := make(chan map[string]string)
	minV := minConfig()
	wg.Add(5)
	go sarm.GetPageUrlList(minV.GetString("国家市场监督管理总局"), urlChannel, wg)
	go miit.GetPageUrlList(minV.GetString("工业和信息化部"), urlChannel, wg)
	go mee.GetFirstUrl(minV.GetString("生态环境部"), urlChannel, wg)
	go cbirc.GetPageUrlList(minV.GetString("银保监会928"), infoChannel, wg)
	go cbirc.GetPageUrlList(minV.GetString("银保监会927"), infoChannel, wg)

	go func() {
		for v := range urlChannel {
			wg.Add(1)
			go v.GetUrlFunc(urlChannel, infoChannel, wg)

		}
	}()
	go func() {
		for v := range infoChannel {
			wg.Add(1)
			go v.GetInfoFunc(errChannel, infoMap, wg)
		}
	}()
	go func() {
		wg.Wait()
		close(urlChannel)
		close(infoChannel)
		close(infoMap)
	}()

	for info := range infoMap {
		for _, k := range info {
			if k != "" {
				fmt.Println(info)
			}
		}
	}
}

func main() {
	r := M.Query("SELECT * FROM t_dmaa_base_target")
	fmt.Println(len(r))
	ministries()
}
