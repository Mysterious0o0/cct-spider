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

// 主要经济指标-全国年度数据-居民人均消费支出(元)
// 主要经济指标-全国年度数据-居民人均消费支出_同比增长(%)

func hce() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	hceRegion := last.YearRegion(indexcode.HCEStartYear)
	for _, region := range hceRegion {
		c := core.Core{
			TL: "year",
			SQL: fmt.Sprintf(sql, indexcode.HCECode),
			IndexCode: indexcode.HCECode,
			TypeCode: typecode.YearDataCode,
			UnitType: "",
			UnitName: "元",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgnd",
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
		time.Sleep(time.Second * 5)
	}
}

func hce1() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	hce1Region := last.YearRegion(indexcode.HCE1StartYear)
	for _, region := range hce1Region {
		c := core.Core{
			TL: "year",
			SQL: fmt.Sprintf(sql, indexcode.HCE1Code),
			IndexCode: indexcode.HCE1Code,
			TypeCode: typecode.YearDataCode,
			UnitType: "",
			UnitName: "%",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgnd",
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
		time.Sleep(time.Second * 5)
	}
}

func runHCENationYear() {
	hce()
	hce1()
}

func Run() {
	runHCENationYear()
}