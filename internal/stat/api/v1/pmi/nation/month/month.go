package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国月度数据-制造业采购经理指数(%)

func runPMINationMonth() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	corePMI := core.Core{
		TL: "month",
		SQL: fmt.Sprintf(sql, indexcode.PMICode),
		IndexCode: indexcode.PMICode,
		TypeCode: typecode.MonthDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgyd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.PMIStartYear, 1),
		},
	}
	corePMI.Run()
}

func Run() {
	runPMINationMonth()
}