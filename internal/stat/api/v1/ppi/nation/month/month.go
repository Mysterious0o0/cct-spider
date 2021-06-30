package month

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上月=100)
// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上年同月=100)
// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上年同期=100)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.PPI2StartYear, 1),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.PPI2Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.PPI3StartYear, 10),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.PPI3Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.PPI4StartYear, 1),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.PPI4Code)
}
