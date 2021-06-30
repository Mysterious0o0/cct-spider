package month

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-社会消费品零售总额_当期值(亿元)
// 主要经济指标-全国月度数据-社会消费品零售总额_累计值(亿元)
// 主要经济指标-全国月度数据-社会消费品零售总额_同比增长(%)
// 主要经济指标-全国月度数据-社会消费品零售总额_累计增长(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.SCG1StartYear, 1),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.SCG1Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.SCG2StartYear, 1),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.SCG2Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.SCG3StartYear, 1),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.SCG3Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.SCG4StartYear, 1),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.SCG4Code)
}
