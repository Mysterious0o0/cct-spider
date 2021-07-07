package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上月=100)
// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上年同月=100)
// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上年同期=100)

func runPPINationMonth() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	corePPI2 := core.Core{
		TL:        "month",
		SQL:       fmt.Sprintf(sql, indexcode.PPI2Code),
		IndexCode: indexcode.PPI2Code,
		TypeCode:  typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.PPI2StartYear, 1),
		},
	}
	corePPI2.Run()

	corePPI3 := core.Core{
		TL:        "month",
		SQL:       fmt.Sprintf(sql, indexcode.PPI3Code),
		IndexCode: indexcode.PPI3Code,
		TypeCode:  typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.PPI3StartYear, 10),
		},
	}
	corePPI3.Run()

	corePPI4 := core.Core{
		TL:        "month",
		SQL:       fmt.Sprintf(sql, indexcode.PPI4Code),
		IndexCode: indexcode.PPI4Code,
		TypeCode:  typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.PPI4StartYear, 1),
		},
	}
	corePPI4.Run()
}

func Run() {
	runPPINationMonth()
}