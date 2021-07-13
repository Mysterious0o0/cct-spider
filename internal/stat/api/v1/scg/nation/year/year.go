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

// 主要经济指标-全国年度数据-社会消费品零售总额(亿元)

func runSCGNationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.SCGName
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	scgRegion := last.YearRegion(start)

	for _, region := range scgRegion {
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
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}
}

func Run() {
	runSCGNationYear()
}