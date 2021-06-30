package year

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国年度数据-城镇居民人均可支配收入(元)
// 主要经济指标-全国年度数据-城镇居民人均可支配收入_同比增长(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgnd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Years(indexcode.URIStartYear),
	}
	executor.Executor(url, typecode.YearDataCode, indexcode.URICode)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgnd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Years(indexcode.URI1StartYear),
	}
	executor.Executor(url, typecode.YearDataCode, indexcode.URI1Code)
}



