package month

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-制造业采购经理指数(%)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgyd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Months(indexcode.PMIStartYear, 1),
	}
	executor.Executor(url, typecode.MonthDataCode, indexcode.PMICode)
}
