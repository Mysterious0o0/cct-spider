package season

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国季度数据-固定资产投资价格指数_当季值(上年同季=100)
// 主要经济指标-全国季度数据-固定资产投资价格指数_累计值(上年同期=100)

func runFaiNationSeason() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreFAI2 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.FAI2Code),
		IndexCode: indexcode.FAI2Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.FAI2StartYear, 2),
		},
	}
	coreFAI2.Run()

	coreFAI3 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.FAI3Code),
		IndexCode: indexcode.FAI3Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.FAI3StartYear, 1),
		},
	}
	coreFAI3.Run()
}

func Run() {
	runFaiNationSeason()
}
