package season

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-工业增加值_当季值(亿元)
// 主要经济指标-全国月度数据-工业增加值_累计值(亿元)

func runIAVNationSeason() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreIAV3 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.IAV3Code),
		IndexCode: indexcode.IAV3Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.IAV3StartYear, 1),
		},
	}
	coreIAV3.Run()

	coreIAV4 := core.Core{
		TL: "season",
		SQL: fmt.Sprintf(sql, indexcode.IAV4Code),
		IndexCode: indexcode.IAV4Code,
		TypeCode: typecode.SeasonDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgjd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Seasons(indexcode.IAV4StartYear, 1),
		},
	}
	coreIAV4.Run()
}

func Run() {
	runIAVNationSeason()
}