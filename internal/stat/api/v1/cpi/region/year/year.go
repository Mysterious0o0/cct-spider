package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"github.com/xiaogogonuo/cct-spider/pkg/set"
)

// 地区经济指标-分省年度数据-居民消费价格指数(上年=100)

func runCPI(indexCode, typeCode string, startYear int) {
	for _, v := range provincecode.ProvinceCode {
		url := urllib.Param{
			M:              "QueryData",
			DBCode:         "fsnd",
			RowCode:        "zb",
			ColCode:        "sj",
			WdsWdCode:      "reg",
			WdsWdValueCode: v,  // 省、市编码
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(startYear),
		}
		querySQL := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s' AND REGION_CODE = '%s'`
		queryRow := mysql.Query(fmt.Sprintf(querySQL, indexCode, v))
		row := executor.Executor(url, typeCode, indexCode)
		diff, err := set.Set{Src: queryRow}.Diff(row)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		dataBuilder := response.DataBuilder{
			IndexCode: indexCode,
			RegionCode: v,
			IsYear: true,
		}
		data := dataBuilder.Build(diff)
		if len(data) == 0 {
			logger.Info(fmt.Sprintf("%s has no new data to update",indexcode.CodeName[indexCode]))
			continue
		}
		tranSQL := mysql.Generator(data)
		mysql.Transaction(tranSQL)
	}
}

func runCPIRegionYear() {
	runCPI(indexcode.CPICode, typecode.ProvinceYearDataCode, indexcode.CPIStartYear)

	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s' AND REGION_CODE = '%s'`

	for _, v := range provincecode.ProvinceCode {
		coreCPI := core.Core{
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
				DfWdsValueCode: last.Years(indexcode.CPIStartYear),
			},
		}
		coreCPI.Run()
	}
}

func Run() {
	runCPIRegionYear()
}