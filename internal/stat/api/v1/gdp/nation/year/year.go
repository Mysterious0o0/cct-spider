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

// 主要经济指标-全国年度数据-国内生产总值年度(亿元)

func runGDPNationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.GDPName
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	gdpRegion := last.YearRegion(start)

	for _, region := range gdpRegion {
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
			UnitName: "亿元",
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
		rowsAffected, err :=c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 5)
	}
}

func Run() {
	runGDPNationYear()
}