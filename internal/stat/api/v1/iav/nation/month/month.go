package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-工业增加值_同比增长(%)
// 主要经济指标-全国月度数据-工业增加值_累计增长(%)

func runIAVNationMonth() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreIAV1 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.IAV1Code),
		IndexCode: indexcode.IAV1Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.IAV1StartYear, 7),
		},
	}
	coreIAV1.Run()

	coreIAV2 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.IAV2Code),
		IndexCode: indexcode.IAV2Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.IAV2StartYear, 7),
		},
	}
	coreIAV2.Run()
}

func Run() {
	runIAVNationMonth()
}