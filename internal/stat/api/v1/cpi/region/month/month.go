package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"strconv"
	"time"
)

// 地区经济指标-分省月度数据-居民消费价格指数(上年同月=100)[start: 2016A]

func runCPIRegionMonth() {
	sql := `SELECT CONCAT(ACCT_YEAR, ACCT_MONTH), TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s' AND REGION_CODE = '%s'`

	indexName := indexcode.CPI3Name
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	cpi3Region := last.YearRegion(start)

	for _, region := range cpi3Region {
		for _, v := range provincecode.ProvinceCode {
			c := core.Core{
				TL: "month",
				SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"], v),
				IndexName: indexName,
				TypeCode: typecode.ProvinceMonthDataCode,
				URL: urllib.Param{
					M:              "QueryData",
					DBCode:         "fsyd",
					RowCode:        "zb",
					ColCode:        "sj",
					WdsWdCode:      "reg",
					WdsWdValueCode: v,
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
}

func Run() {
	runCPIRegionMonth()
}