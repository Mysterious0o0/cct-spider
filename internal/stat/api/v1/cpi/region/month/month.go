package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 地区经济指标-分省月度数据-居民消费价格指数(上年同月=100)

func runCPIRegionMonth() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s' AND REGION_CODE = '%s'`

	for _, v := range provincecode.ProvinceCode {
		coreCPI3 := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.CPI3Code, v),
			IndexCode: indexcode.CPI3Code,
			TypeCode: typecode.ProvinceMonthDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "fsyd",
				RowCode:        "zb",
				ColCode:        "sj",
				WdsWdCode:      "reg",
				WdsWdValueCode: v,
				DfWdsWdCode:    "sj",
				DfWdsValueCode: last.Months(indexcode.CPI3StartYear, 1),
			},
		}
		coreCPI3.Run()
	}
}

func Run() {
	runCPIRegionMonth()
}