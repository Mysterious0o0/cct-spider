package season

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国季度数据-城镇居民人均可支配收入_累计值(元)
// 主要经济指标-全国季度数据-城镇居民人均可支配收入_累计增长(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.URI2StartYear, 1),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.URI2Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.URI3StartYear, 1),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.URI3Code)
}
