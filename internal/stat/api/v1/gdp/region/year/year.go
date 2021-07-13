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
	"time"
)

// 地区经济指标-分省年度数据-地区生产总值(亿元)

func runGDPRegionYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s' AND REGION_CODE = '%s'`

	indexName := indexcode.GDPRName
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	gdpRRegion := last.YearRegion(start)

	for _, v := range provincecode.ProvinceCode {
		for _, region := range gdpRRegion {
			c := core.Core{
				TL: "year",
				SQL: fmt.Sprintf(sql, indexcode.IndexMap[indexName]["innerCode"], v),
				IndexName: indexName,
				TypeCode: typecode.ProvinceYearDataCode,
				UnitType: "",
				UnitName: "亿元",
				URL: urllib.Param{
					M:              "QueryData",
					DBCode:         "fsnd",
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
				break
			}
			time.Sleep(time.Second * 5)
		}
	}
}

func Run() {
	runGDPRegionYear()
}