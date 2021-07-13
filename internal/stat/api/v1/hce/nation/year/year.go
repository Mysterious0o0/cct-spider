package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"strconv"
	"strings"
	"time"
)

// 主要经济指标-全国年度数据-居民人均消费支出(元)
// 主要经济指标-全国年度数据-居民人均消费支出_同比增长(%)

func hce() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.HCEName
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	hceRegion := last.YearRegion(start)

	for _, region := range hceRegion {
		// 过滤掉今年
		stop, _ := strconv.Atoi(strings.Split(region, "-")[1])
		if stop == time.Now().Year() {
			continue
		}
		c := core.Core{
			TL: "year",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
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
			IndexMap: indexcode.IndexMap,
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func hce1() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.HCE1Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	hce1Region := last.YearRegion(start)

	for _, region := range hce1Region {
		// 过滤掉今年
		stop, _ := strconv.Atoi(strings.Split(region, "-")[1])
		if stop == time.Now().Year() {
			continue
		}
		c := core.Core{
			TL: "year",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
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
			IndexMap: indexcode.IndexMap,
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func runHCENationYear() {
	hce()
	hce1()
}

func Run() {
	runHCENationYear()
}