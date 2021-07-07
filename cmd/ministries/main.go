package main

import (
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/cbirc"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/mee"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/miit"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/sarm"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"sync"
)

var (
	minV *viper.Viper
	filter *store.Filter
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

func init() {
	minV = minConfig()
	filter = &store.Filter{
		Filepath: "urlKey.txt",
	}
	filter.ReadUrlKey()
}

func ministries() {
	wg := &sync.WaitGroup{}
	urlChannel := make(chan *store.UrlChan)   // url请求池
	infoChannel := make(chan *store.InfoChan) // info请求池
	errChannel := make(chan *store.InfoChan)  // 异常池
	message := make(chan *store.Message)      // 数据池
	save := store.InsertIntoSQL               // 保存数据的函数

	wg.Add(5)
	go miit.GetPageUrlList(minV.GetString("工业和信息化部"), urlChannel, wg)
	go sarm.GetPageUrlList(minV.GetString("国家市场监督管理总局"), urlChannel, wg)
	go mee.GetFirstUrl(minV.GetString("生态环境部"), urlChannel, wg)
	go cbirc.GetPageUrlList(minV.GetString("银保监会928"), infoChannel, wg)
	go cbirc.GetPageUrlList(minV.GetString("银保监会927"), infoChannel, wg)

	go func() {
		for v := range urlChannel {
			if _, ok := filter.UrlKey[md5.MD5(v.Url)]; ok{
				continue
			}
			wg.Add(1)
			go v.GetUrlFunc(urlChannel, infoChannel, wg)
		}
	}()
	go func() {
		for v := range infoChannel {
			if _, ok := filter.UrlKey[md5.MD5(v.Url)]; ok{
				continue
			}
			wg.Add(1)
			go v.GetInfoFunc(errChannel, message, wg)
		}
	}()
	go func() {
		wg.Wait()
		close(urlChannel)
		close(infoChannel)
		close(message)
	}()
	save(filter, message)

}

func main() {
	ministries()
}
