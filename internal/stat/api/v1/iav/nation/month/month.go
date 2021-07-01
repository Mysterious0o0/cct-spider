package month

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-工业增加值_同比增长(%)
// 主要经济指标-全国月度数据-工业增加值_累计增长(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.IAV1StartYear, 7),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.IAV1Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.IAV2StartYear, 7),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.IAV2Code)
}
