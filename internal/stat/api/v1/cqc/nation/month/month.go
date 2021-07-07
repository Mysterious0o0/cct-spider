package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-货币和准货币(M2)供应量_期末值(亿元)
// 主要经济指标-全国月度数据-货币和准货币(M2)供应量_同比增长(%)

func runCQCNationMonth() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreCQC1 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.CQC1Code),
		IndexCode: indexcode.CQC1Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd", // 宏观月度
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.CQC1StartYear, 12),
		},
	}
	coreCQC1.Run()

	coreCQC2 := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.CQC2Code),
		IndexCode: indexcode.CQC2Code,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd", // 宏观月度
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.CQC2StartYear, 12),
		},
	}
	coreCQC2.Run()
}

func Run() {
	runCQCNationMonth()
}