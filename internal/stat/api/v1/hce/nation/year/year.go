package year

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国年度数据-居民人均消费支出(元)
// 主要经济指标-全国年度数据-居民人均消费支出_同比增长(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgnd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Years(indexcode.HCEStartYear),
	}
	executor.Executor(url, typecode.YearDataCode, indexcode.HCECode)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgnd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Years(indexcode.HCE1StartYear),
	}
	executor.Executor(url, typecode.YearDataCode, indexcode.HCE1Code)
}
