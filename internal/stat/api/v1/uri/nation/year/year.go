package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"time"
)

// 主要经济指标-全国年度数据-城镇居民人均可支配收入(元)
// 主要经济指标-全国年度数据-城镇居民人均可支配收入_同比增长(%)

func uri() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	uriRegion := last.YearRegion(indexcode.URIStartYear)
	for _, region := range uriRegion {
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
				DfWdsValueCode: region,
			},
		}
		rowsAffected, err := coreURI.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func uri1() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	uri1Region := last.YearRegion(indexcode.URI1StartYear)
	for _, region := range uri1Region {
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
				DfWdsValueCode: region,
			},
		}
		rowsAffected, err := coreURI1.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func runURINationYear() {
	uri()
	uri1()
}

func Run() {
	runURINationYear()
}