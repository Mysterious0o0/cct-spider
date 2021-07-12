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

// 主要经济指标-全国月度数据-货币和准货币(M2)供应量_期末值(亿元)
// 主要经济指标-全国月度数据-货币和准货币(M2)供应量_同比增长(%)

func cqc1() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	cqc1Region := last.YearRegion(indexcode.CQC1StartYear)
	for _, region := range cqc1Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.CQC1Code),
			IndexCode: indexcode.CQC1Code,
			TypeCode: typecode.MonthDataCode,
			UnitType: "",
			UnitName: "亿元",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgyd", // 宏观月度
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

func cqc2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	cqc2Region := last.YearRegion(indexcode.CQC2StartYear)
	for _, region := range cqc2Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.CQC2Code),
			IndexCode: indexcode.CQC2Code,
			TypeCode: typecode.MonthDataCode,
			UnitType: "",
			UnitName: "%",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgyd", // 宏观月度
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

func runCQCNationMonth() {
	cqc1()
	cqc2()
}

func Run() {
	runCQCNationMonth()
}