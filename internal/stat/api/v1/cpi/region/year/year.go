package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"time"
)

// 地区经济指标-分省年度数据-居民消费价格指数(上年=100)

func runCPIRegionYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s' AND REGION_CODE = '%s'`

	cpiRegion := last.YearRegion(indexcode.CPIStartYear)
	for _, region := range cpiRegion {
		for _, v := range provincecode.ProvinceCode {
			c := core.Core{
				TL: "year",
				SQL: fmt.Sprintf(sql, indexcode.CPICode, v),
				IndexCode: indexcode.CPICode,
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