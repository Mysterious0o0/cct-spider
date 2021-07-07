package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 地区经济指标-分省年度数据-居民消费水平(元)

func runRCLRegionYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE
              WHERE SOURCE_TARGET_CODE = '%s' AND REGION_CODE = '%s'`

	for _, v := range provincecode.ProvinceCode {
		coreRCL := core.Core{
			TL: "year",
			SQL: fmt.Sprintf(sql, indexcode.RCLCode, v),
			IndexCode: indexcode.RCLCode,
			TypeCode: typecode.ProvinceYearDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "fsnd",
				RowCode:        "zb",
				ColCode:        "sj",
				WdsWdCode:      "reg",
				WdsWdValueCode: v,
				DfWdsWdCode:    "sj",
				DfWdsValueCode: last.Years(indexcode.RCLStartYear),
			},
		}
		coreRCL.Run()
	}
}

func Run() {
	runRCLRegionYear()
}