package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"strconv"
	"time"
)

// 主要经济指标-全国月度数据-工业增加值_同比增长(%)
// 主要经济指标-全国月度数据-工业增加值_累计增长(%)

func iav1() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.IAV1Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	iav1Region := last.YearRegion(start)

	for _, region := range iav1Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.MonthDataCode,
			UnitType: "",
			UnitName: "%",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgyd",
				RowCode:        "zb",
				ColCode:        "sj",
				DfWdsWdCode:    "sj",
				DfWdsValueCode: region,
			},
			IndexMap: indexcode.IndexMap,
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func iav2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.IAV2Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	iav2Region := last.YearRegion(start)

	for _, region := range iav2Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.MonthDataCode,
			UnitType: "",
			UnitName: "%",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgyd",
				RowCode:        "zb",
				ColCode:        "sj",
				DfWdsWdCode:    "sj",
				DfWdsValueCode: region,
			},
			IndexMap: indexcode.IndexMap,
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func runIAVNationMonth() {
	iav1()
	iav2()
}

func Run() {
	runIAVNationMonth()
}