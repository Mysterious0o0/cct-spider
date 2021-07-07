package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国年度数据-社会消费品零售总额(亿元)

func runSCGNationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreSCG := core.Core{
		TL: "year",
		SQL: fmt.Sprintf(sql, indexcode.SCGCode),
		IndexCode: indexcode.SCGCode,
		TypeCode: typecode.YearDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgnd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(indexcode.SCGStartYear),
		},
	}
	coreSCG.Run()
}

func Run() {
	runSCGNationYear()
}