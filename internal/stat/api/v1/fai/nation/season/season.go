package season

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国季度数据-固定资产投资价格指数_当季值(上年同季=100)
// 主要经济指标-全国季度数据-固定资产投资价格指数_累计值(上年同期=100)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.FAI2StartYear, 2),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.FAI2Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.FAI3StartYear, 1),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.FAI3Code)
}
