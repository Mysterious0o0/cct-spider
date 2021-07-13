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

// 主要经济指标-全国季度数据-城镇居民人均可支配收入_累计值(元)
// 主要经济指标-全国季度数据-城镇居民人均可支配收入_累计增长(%)

func uri2() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.URI2Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	uri2Region := last.YearRegion(start)

	for _, region := range uri2Region {
		c := core.Core{
			TL: "season",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.SeasonDataCode,
			UnitType: "",
			UnitName: "元",
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

func uri3() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.URI3Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	uri3Region := last.YearRegion(start)

	for _, region := range uri3Region {
		c := core.Core{
			TL: "season",
			SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"]),
			IndexName: indexName,
			TypeCode: typecode.SeasonDataCode,
			UnitType: "",
			UnitName: "%",
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

func runURINationSeason() {
	uri2()
	uri3()
}

func Run() {
	runURINationSeason()
}