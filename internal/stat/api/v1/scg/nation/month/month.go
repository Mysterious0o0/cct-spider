package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-社会消费品零售总额_当期值(亿元)
// 主要经济指标-全国月度数据-社会消费品零售总额_累计值(亿元)
// 主要经济指标-全国月度数据-社会消费品零售总额_同比增长(%)
// 主要经济指标-全国月度数据-社会消费品零售总额_累计增长(%)

func runSCGNationMonth() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreSCG1 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.SCG1Code),
		IndexCode: indexcode.SCG1Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.SCG1StartYear, 1),
		},
	}
	coreSCG1.Run()

	coreSCG2 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.SCG2Code),
		IndexCode: indexcode.SCG2Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.SCG2StartYear, 1),
		},
	}
	coreSCG2.Run()

	coreSCG3 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.SCG3Code),
		IndexCode: indexcode.SCG3Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.SCG3StartYear, 1),
		},
	}
	coreSCG3.Run()

	coreSCG4 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.SCG4Code),
		IndexCode: indexcode.SCG4Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.SCG4StartYear, 1),
		},
	}
	coreSCG4.Run()
}

func Run() {
	runSCGNationMonth()
}