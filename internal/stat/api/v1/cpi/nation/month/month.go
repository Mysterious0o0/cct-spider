package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"time"
)

// 主要经济指标-全国月度数据-居民消费价格指数(上月=100)
// 主要经济指标-全国月度数据-居民消费价格指数(上年同月=100)
// 主要经济指标-全国月度数据-居民消费价格指数(上年同期=100)

func cpi2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	cpi2Region := last.YearRegion(indexcode.CPI2StartYear)
	for _, region := range cpi2Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.CPI2Code),
			IndexCode: indexcode.CPI2Code,
			TypeCode: typecode.MonthDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgyd",
				RowCode:        "zb",
				ColCode:        "sj",
				DfWdsWdCode:    "sj",
				DfWdsValueCode: region,
			},
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func cpi3() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	cpi3Region := last.YearRegion(indexcode.CPI3StartYear)
	for _, region := range cpi3Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.CPI3Code),
			IndexCode: indexcode.CPI3Code,
			TypeCode: typecode.MonthDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgyd",
				RowCode:        "zb",
				ColCode:        "sj",
				DfWdsWdCode:    "sj",
				DfWdsValueCode: region,
			},
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func cpi4() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	cpi4Region := last.YearRegion(indexcode.CPI4StartYear)
	for _, region := range cpi4Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.CPI4Code),
			IndexCode: indexcode.CPI4Code,
			TypeCode: typecode.MonthDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgyd",
				RowCode:        "zb",
				ColCode:        "sj",
				DfWdsWdCode:    "sj",
				DfWdsValueCode: region,
			},
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func runCPINationMonth() {
	cpi2()
    cpi3()
    cpi4()
}

func Run() {
	runCPINationMonth()
}