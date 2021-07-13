package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"strconv"
	"strings"
	"time"
)

// 地区经济指标-分省年度数据-居民消费价格指数(上年=100)

func runCPIRegionYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s' AND REGION_CODE = '%s'`

	indexName := indexcode.CPIName
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	cpiRegion := last.YearRegion(start)

	for _, region := range cpiRegion {
		// 过滤掉今年
		stop, _ := strconv.Atoi(strings.Split(region, "-")[1])
		if stop == time.Now().Year() {
			continue
		}
		for _, v := range provincecode.ProvinceCode {
			c := core.Core{
				TL: "year",
				SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"], v),
				IndexName: indexName,
				TypeCode: typecode.ProvinceYearDataCode,
				URL: urllib.Param{
					M:              "QueryData",
					DBCode:         "fsnd",
					RowCode:        "zb",
					ColCode:        "sj",
					WdsWdCode:      "reg",
					WdsWdValueCode: v,  // 省、市编码
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
	runCPIRegionYear()
}