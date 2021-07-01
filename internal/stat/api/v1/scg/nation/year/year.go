package year

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国年度数据-社会消费品零售总额(亿元)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgnd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Years(indexcode.SCGStartYear),
	}
	executor.Executor(url, typecode.YearDataCode, indexcode.SCGCode)
}
