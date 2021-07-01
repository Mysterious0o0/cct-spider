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
	sql := `
INSERT INTO  t_dmaa_base_target_value
    (
     VALUE_GUID,
     TARGET_GUID,
     TARGET_CODE,
     TARGET_NAME,
     SOURCE_TARGET_CODE,
     REGION_CODE,
     UNIT_TYPE,
     UNIT_NAME,
     ACCT_YEAR,
     ACCT_QUARTOR,
     ACCT_MONTH,
     TARGET_VALUE
     )
VALUES
       ('1', '2c20e4831b1f334fd58d685e32de2f8c', 'HG00001', '国内生产总值', 'A020102', '', '人民币' ,'亿元', '2021', '', '', '102'),
       ('2', '19dfc1ebad7e4c9c2cb39fd884bbe9dc', 'HG00002', '地区生产总值', 'A020101', '310000', '人民币' ,'亿元', '2021', '', '6', '5000')
ON DUPLICATE KEY UPDATE
    VALUE_GUID = VALUES(VALUE_GUID),
    TARGET_GUID = VALUES(TARGET_GUID),
    TARGET_CODE = VALUES(TARGET_CODE),
    TARGET_NAME = VALUES(TARGET_NAME),
    SOURCE_TARGET_CODE = VALUES(SOURCE_TARGET_CODE),
    REGION_CODE = VALUES(REGION_CODE),
    UNIT_TYPE = VALUES(UNIT_TYPE),
    UNIT_NAME = VALUES(UNIT_NAME),
    ACCT_YEAR = VALUES(ACCT_YEAR),
    ACCT_QUARTOR = VALUES(ACCT_QUARTOR),
    ACCT_MONTH = VALUES(ACCT_MONTH),
    TARGET_VALUE = VALUES(TARGET_VALUE);
`
	M.Transaction(sql)
}
