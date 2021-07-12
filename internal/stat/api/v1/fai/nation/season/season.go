package season

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"time"
)

// 主要经济指标-全国季度数据-固定资产投资价格指数_当季值(上年同季=100)
// 主要经济指标-全国季度数据-固定资产投资价格指数_累计值(上年同期=100)

func fai2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	fai2Region := last.YearRegion(indexcode.FAI2StartYear)
	for _, region := range fai2Region {
		c := core.Core{
			TL: "season",
			SQL: fmt.Sprintf(sql, indexcode.FAI2Code),
			IndexCode: indexcode.FAI2Code,
			TypeCode: typecode.SeasonDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgjd",
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

func fai3() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	fai3Region := last.YearRegion(indexcode.FAI3StartYear)
	for _, region := range fai3Region {
		c := core.Core{
			TL: "season",
			SQL: fmt.Sprintf(sql, indexcode.FAI3Code),
			IndexCode: indexcode.FAI3Code,
			TypeCode: typecode.SeasonDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgjd",
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

func runFaiNationSeason() {
	fai2()
	fai3()
}

func Run() {
	runFaiNationSeason()
}
