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

// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上月=100)
// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上年同月=100)
// 主要经济指标-全国月度数据-工业生产者出厂价格指数(上年同期=100)

func ppi2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	ppi2Region := last.YearRegion(indexcode.PPI2StartYear)
	for _, region := range ppi2Region {
		c := core.Core{
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

func ppi3() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	ppi3Region := last.YearRegion(indexcode.PPI3StartYear)
	for _, region := range ppi3Region {
		c := core.Core{
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

func ppi4() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	ppi4Region := last.YearRegion(indexcode.PPI4StartYear)
	for _, region := range ppi4Region {
		c := core.Core{
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

func runPPINationMonth() {
	ppi2()
	ppi3()
	ppi4()
}

func Run() {
	runPPINationMonth()
}