package month

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-货币和准货币(M2)供应量_期末值(亿元)
// 主要经济指标-全国月度数据-货币和准货币(M2)供应量_同比增长(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.CQC1StartYear, 12),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.CQC1Code)

	url = urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.CQC2StartYear, 12),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.CQC2Code)
}
