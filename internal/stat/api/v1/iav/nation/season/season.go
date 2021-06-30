package season

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-工业增加值_当季值(亿元)
// 主要经济指标-全国月度数据-工业增加值_累计值(亿元)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.IAV3StartYear, 1),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.IAV3Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgjd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Seasons(indexcode.IAV4StartYear, 1),
	}
	executor.Executor(url, typecode.SeasonDataCode, indexcode.IAV4Code)
}
