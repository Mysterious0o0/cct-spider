package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国年度数据-居民人均消费支出(元)
// 主要经济指标-全国年度数据-居民人均消费支出_同比增长(%)

func runHCENationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreHCE := core.Core{
		TL: "year",
		SQL: fmt.Sprintf(sql, indexcode.HCECode),
		IndexCode: indexcode.HCECode,
		TypeCode: typecode.YearDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgnd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(indexcode.HCEStartYear),
		},
	}
	coreHCE.Run()

	coreHCE1 := core.Core{
		TL: "year",
		SQL: fmt.Sprintf(sql, indexcode.HCE1Code),
		IndexCode: indexcode.HCE1Code,
		TypeCode: typecode.YearDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgnd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(indexcode.HCE1StartYear),
		},
	}
	coreHCE1.Run()
}

func Run() {
	runHCENationYear()
}