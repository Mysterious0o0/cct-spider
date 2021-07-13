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

// 主要经济指标-全国月度数据-社会消费品零售总额_当期值(亿元)
// 主要经济指标-全国月度数据-社会消费品零售总额_累计值(亿元)
// 主要经济指标-全国月度数据-社会消费品零售总额_同比增长(%)
// 主要经济指标-全国月度数据-社会消费品零售总额_累计增长(%)

func scg1() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.SCG1Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	scg1Region := last.YearRegion(start)

	for _, region := range scg1Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.MonthDataCode,
			UnitType: "",
			UnitName: "亿元",
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

func scg2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.SCG2Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	scg2Region := last.YearRegion(start)

	for _, region := range scg2Region {
		c := core.Core{
			TL: "month",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.MonthDataCode,
			UnitType: "",
			UnitName: "亿元",
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

func scg3() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.SCG3Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	scg3Region := last.YearRegion(start)

	for _, region := range scg3Region {
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

func scg4() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.SCG4Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	scg4Region := last.YearRegion(start)

	for _, region := range scg4Region {
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

func runSCGNationMonth() {
	scg1()
	scg2()
	scg3()
	scg4()
}

func Run() {
	runSCGNationMonth()
}