package season

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

// 主要经济指标-全国季度数据-国内生产总值_当季值(亿元)[start: 1992A]
// 主要经济指标-全国季度数据-国内生产总值_累计值(亿元)[start: 1992A]

func gdp1() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.GDP1Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	gdp1Region := last.YearRegion(start)

	for _, region := range gdp1Region {
		c := core.Core{
			TL: "season",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.SeasonDataCode,
			UnitType: "",
			UnitName: "亿元",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgjd",
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

func gdp2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.GDP2Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	gdp2Region := last.YearRegion(start)

	for _, region := range gdp2Region {
		c := core.Core{
			TL: "season",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.SeasonDataCode,
			UnitType: "",
			UnitName: "亿元",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgjd",
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

func runGDPNationSeason() {
	gdp1()
	gdp2()
}

func Run() {
	runGDPNationSeason()
}