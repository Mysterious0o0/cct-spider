package year

import (
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国年度数据-工业生产者出厂价格指数(上年=100)
// 主要经济指标-全国年度数据-工业生产者出厂价格指数(1985=100)

func Run() {
	url := urllib.Param{
		M:              "QueryData",
		DBCode:         "hgnd",
		RowCode:        "zb",
		ColCode:        "sj",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: last.Years(indexcode.PPIStartYear),
	}
	executor.Executor(url, typecode.YearDataCode, indexcode.PPICode)
}
