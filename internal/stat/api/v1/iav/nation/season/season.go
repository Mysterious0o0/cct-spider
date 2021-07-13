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

// 主要经济指标-全国月度数据-工业增加值_当季值(亿元)
// 主要经济指标-全国月度数据-工业增加值_累计值(亿元)

func iav3() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.IAV3Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	iav3Region := last.YearRegion(start)

	for _, region := range iav3Region {
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

func iav4() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_QUARTOR), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.IAV4Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	iav4Region := last.YearRegion(start)

	for _, region := range iav4Region {
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

func runIAVNationSeason() {
	iav3()
	iav4()
}

func Run() {
	runIAVNationSeason()
}