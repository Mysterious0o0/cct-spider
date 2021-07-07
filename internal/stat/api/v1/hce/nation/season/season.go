package season

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国季度数据-居民人均消费支出_累计值(元)
// 主要经济指标-全国季度数据-居民人均消费支出_累计增长(%)

func runHCENationSeason() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreHCE2 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.HCE2Code),
		IndexCode: indexcode.HCE2Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.HCE2StartYear, 1),
		},
	}
	coreHCE2.Run()

	coreHCE3 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.HCE3Code),
		IndexCode: indexcode.HCE3Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.HCE3StartYear, 1),
		},
	}
	coreHCE3.Run()
}

func Run() {
	runHCENationSeason()
}