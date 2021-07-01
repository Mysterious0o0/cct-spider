package year

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"time"
)

// 地区经济指标-分省年度数据-地区生产总值(亿元)

func Run() {
	for _, v := range provincecode.ProvinceCode {
		url := urllib.Param{
			M:              "QueryData",
			DBCode:         "fsnd",
			RowCode:        "zb",
			ColCode:        "sj",
			WdsWdCode:      "reg",
			WdsWdValueCode: v,
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(indexcode.GDPRStartYear),
		}
		executor.Executor(url, typecode.ProvinceYearDataCode, indexcode.GDPRCode)
		time.Sleep(time.Second * 3)
	}
}