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

// 主要经济指标-全国年度数据-货币和准货币(M2)供应量同比增长率(%)

func runCQCNationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE TARGET_CODE = '%s'`

	indexName := indexcode.GDPName
	startYear := indexcode.IndexMap[indexName]["startYear"]
	start, _ := strconv.Atoi(startYear)
	cqcRegion := last.YearRegion(start)

	for _, region := range cqcRegion {
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
				DBCode:         "hgnd", // 宏观年度
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
	runCQCNationYear()
}