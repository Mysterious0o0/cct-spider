package season

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国季度数据-国内生产总值_当季值(亿元)
// 主要经济指标-全国季度数据-国内生产总值_累计值(亿元)

func runGDPNationSeason() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreGDP1 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.GDP1Code),
		IndexCode: indexcode.GDP1Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.GDP1StartYear, 1),
		},
	}
	coreGDP1.Run()

	coreGDP2 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.GDP2Code),
		IndexCode: indexcode.GDP2Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.GDP2StartYear, 1),
		},
	}
	coreGDP2.Run()
}

func Run() {
	runGDPNationSeason()
}