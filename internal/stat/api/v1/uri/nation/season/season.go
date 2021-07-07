package season

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国季度数据-城镇居民人均可支配收入_累计值(元)
// 主要经济指标-全国季度数据-城镇居民人均可支配收入_累计增长(%)

func runURINationSeason() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreURI2 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.URI2Code),
		IndexCode: indexcode.URI2Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.URI2StartYear, 1),
		},
	}
	coreURI2.Run()

	coreURI3 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.URI3Code),
		IndexCode: indexcode.URI3Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.URI3StartYear, 1),
		},
	}
	coreURI3.Run()
}

func Run() {
	runURINationSeason()
}