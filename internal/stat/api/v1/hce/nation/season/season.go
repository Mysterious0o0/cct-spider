package season

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国季度数据-居民人均消费支出_累计值(元)
// 主要经济指标-全国季度数据-居民人均消费支出_累计增长(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.HCE2StartYear, 1),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.HCE2Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.HCE3StartYear, 1),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.HCE3Code)
}
