package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
)

// 主要经济指标-全国年度数据-城镇居民人均可支配收入(元)
// 主要经济指标-全国年度数据-城镇居民人均可支配收入_同比增长(%)

func runURINationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	coreURI := core.Core{
		TL: "year",
		SQL: fmt.Sprintf(sql, indexcode.URICode),
		IndexCode: indexcode.URICode,
		TypeCode: typecode.YearDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgnd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(indexcode.URIStartYear),
		},
	}
	coreURI.Run()

	coreURI1 := core.Core{
		TL: "year",
		SQL: fmt.Sprintf(sql, indexcode.URI1Code),
		IndexCode: indexcode.URI1Code,
		TypeCode: typecode.YearDataCode,
		URL: urllib.Param{
			M:              "QueryData",
			DBCode:         "hgnd",
			RowCode:        "zb",
			ColCode:        "sj",
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(indexcode.URI1StartYear),
		},
	}
	coreURI1.Run()
}

func Run() {
	runURINationYear()
}